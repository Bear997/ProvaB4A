package controllers

import (
	"Bear997/api/db"
	"Bear997/api/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CardRepo struct {
	Db *gorm.DB
}

func Card() *CardRepo {
	db := db.DbConnection()
	db.AutoMigrate(&models.Card{})
	return &CardRepo{Db: db}
}

func (repository *CardRepo) CreateCard(c *gin.Context) {
	var card models.Card
	var mysqlErr *mysql.MySQLError

	c.BindJSON(&card)
	err := models.CreateCard(repository.Db, &card)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1364 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": mysqlErr.Message})
		} else {

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}
	c.JSON(http.StatusOK, card)
}
