package utility

import (
	"encoding/json"
	"fmt"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"net/http"
	"net/mail"
	"reflect"
)

func ValidationStruct(err error, c *gin.Context) {
	fmt.Println("qua ci arrivo")
	fmt.Println(err)
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		msg := fmt.Sprintf("" + e.Field + " must be a " + kindOfData(e.Field).String())
		fmt.Println(msg)

		c.JSON(http.StatusBadRequest, gin.H{"error": "" + e.Field + " must be a " + kindOfData(e.Field).String()})
		return
	} else if errors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errors {
			if fieldErr.Tag() == "required" {
				errMsg := fmt.Sprintf("Field %s is required", fieldErr.Field())
				c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
				return
			}
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return nil

}
func ValidatePassword(psw string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range psw {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}
