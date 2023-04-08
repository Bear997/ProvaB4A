package auth

import (
	"Bear997/api/models"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CreateJwt(user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	fmt.Println(secret)
	// Sign and get the complete encoded token as a string using the secret
	fmt.Println(reflect.TypeOf(os.Getenv("JWT_SECRET")))
	tokenString, err := token.SignedString(secret)
	if err != nil {

		fmt.Println(err)
		fmt.Println("sto dopo in getenv")
		return "", err
	}

	return tokenString, nil
}

func ValidateJwt(tokenString string, c *gin.Context, db *gorm.DB) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	fmt.Println(secret)
	var keyfunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}
	token, err := jwt.Parse(tokenString, keyfunc)
	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizeed"})
		fmt.Println("errore nel parsing del token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		id := fmt.Sprintf("%v", claims["id"])
		err := models.GetUserFromId(db, &user, id)
		if err != nil {
			fmt.Println("errore nel recupero dell'user dato il claim del jwt")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizeed"})
			return
		}
		userId := user.ID
		if id != strconv.FormatUint(uint64(userId), 10) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizeed"})
			return
		}
	} else {
		fmt.Println("ston in else, anuothorized")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

}

func GetIdFromToken(tokenString string) string {
	secret := []byte(os.Getenv("JWT_SECRET"))
	fmt.Println(secret)
	var keyfunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}
	token, err := jwt.Parse(tokenString, keyfunc)
	if err != nil {
		fmt.Println("errore nel parsing del token")
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	var id string
	if ok {
		id = fmt.Sprintf("%v", claims["id"])
	}
	return id
}
