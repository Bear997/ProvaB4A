package controllers

import (
	"Bear997/api/db"
	"Bear997/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CardRepo struct {
	Db *gorm.DB
}
type CustomErr struct {
	Status  int    `json:"status"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Card() *CardRepo {
	db := db.DbConnection()
	db.AutoMigrate(&models.Card{})
	return &CardRepo{Db: db}
}

func (repository *CardRepo) CreateCard(c *gin.Context) {
	var card models.Card
	var mysqlErr *mysql.MySQLError

	if err := c.ShouldBindJSON(&card); err != nil {
		if e, ok := err.(*json.UnmarshalTypeError); ok {

			c.JSON(http.StatusBadRequest, gin.H{"error": "" + e.Field + "must be a" + kindOfData(e.Field).String()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println(reflect.TypeOf(card.Title))
	err := models.CreateCard(repository.Db, &card)
	if err != nil {

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1364 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": mysqlErr.Message})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	c.JSON(http.StatusOK, card)
}

func (repository *CardRepo) GetCardFromPosition(c *gin.Context) {

	queryParams := c.Request.URL.Query()
	lat := queryParams.Get("lat")
	lon := queryParams.Get("lon")
	var card models.Card
	var mysqlErr *mysql.MySQLError

	err := models.GetCardFromPosition(repository.Db, &card, lat, lon)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Card not found"})
			return
		}
		if errors.As(err, &mysqlErr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": mysqlErr.Message})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	c.JSON(http.StatusOK, card)
}
func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
