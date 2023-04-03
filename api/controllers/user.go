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
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func User() *UserRepo {
	db := db.DbConnection()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utility.ValidationStruct(err, c)
	}
	hashPsw, err := utility.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error during hash password"})
	}
	user.Password = hashPsw
	models.CreateUser(repository.Db, &user)
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
	c.JSON(http.StatusOK, gin.H{"accessToken:": tokenJwt})
}
