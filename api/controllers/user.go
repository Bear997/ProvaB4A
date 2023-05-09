package controllers

import (
	"Bear997/api/auth"
	"Bear997/api/db"
	"Bear997/api/models"
	"Bear997/api/utility"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func User() *UserRepo {
	db := db.DbConnection()
	db.AutoMigrate(&models.User{}, &models.Card{}, &models.UserCard{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	var mysqlErr *mysql.MySQLError
	if err := c.ShouldBindJSON(&user); err != nil {
		utility.ValidationStruct(err, c)
		return
	}

	err := utility.ValidateEmail(user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid format for email field"})
		return
	}

	if !utility.ValidatePassword(user.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "password must have at least one upper case letter, at least one lower case letter, at least one digit, at least one special character, at least eight characters long."})
		return
	}

	hashPsw, err := utility.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error during hash password"})
	}
	user.Password = hashPsw
	err = models.CreateUser(repository.Db, &user)
	if err != nil {

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "this email is already used"})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	tokenJwt, err := auth.CreateJwt(user)
	c.JSON(http.StatusOK, gin.H{"accessToken": tokenJwt})

}
func (repository *UserRepo) ChangeProfile(c *gin.Context) {

	var user models.User
	tokenString := c.GetHeader("Authorization")
	userId := auth.GetIdFromToken(tokenString)
	err := models.GetUserFromId(repository.Db, &user, userId)
	if err != nil {
		fmt.Println("error user in changeprofile")
		return
	}
	err = models.ChangeProfile(repository.Db, &user)
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) GetUserById(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	err := models.GetUserFromId(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
	}
	c.JSON(http.StatusOK, user)

}

func (repository *UserRepo) Login(c *gin.Context) {
	var user models.User
	type BodyLogin struct {
		Email    string
		Password string
	}
	var body BodyLogin
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "campi compilati in modo errato"})
		return
	}

	err := models.Login(repository.Db, &user, body.Email)
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Email or password wrong"})
			return
		}
	}

	if !utility.CheckPasswordHash(body.Password, user.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	tokenJwt, err := auth.CreateJwt(user)
	c.JSON(http.StatusOK, gin.H{"accessToken": tokenJwt})
}

func (repository *UserRepo) GetUserProfile(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	var user models.User
	userId := auth.GetIdFromToken(tokenString)
	err := models.GetUserFromId(repository.Db, &user, userId)
	if err != nil {
		fmt.Println("error user in userprofile")
	}

	c.JSON(http.StatusOK, user)
}
func (repository *UserRepo) GetAllCardsOfUser(c *gin.Context) {
	//Todo gestione errori
	tokenString := c.GetHeader("Authorization")
	var user models.User
	userId := auth.GetIdFromToken(tokenString)
	var cards []models.Card
	err := models.GetUserFromId(repository.Db, &user, userId)
	if err != nil {
		fmt.Println("error user in assign")
	}
	models.GetAllCardsOfUser(repository.Db, userId, &user, &cards)
	for _, card := range cards {
		fmt.Println(card.ID)
	}
	c.JSON(http.StatusOK, cards)
}
