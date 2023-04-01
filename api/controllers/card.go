package controllers

import (
	"Bear997/api/db"
	"Bear997/api/models"
	"Bear997/api/utility"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	image, errImage := c.FormFile("image")
	imagePath := "tmp/" + image.Filename
	jsonData := c.Request.Form

	if errImage != nil {
		fmt.Println("sto nell errore dellimagine")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "non riesco a prendere l'immagine"})
		return
	}
	// if err := c.Bind(&card); err != nil {
	// 	utility.ValidationStruct(err, c)
	// 	return
	// }
	card.Title = jsonData.Get("title")
	card.Image = "https://firebasestorage.googleapis.com/v0/b/tearcard-85619.appspot.com/o/" + url.QueryEscape(imagePath) + "?alt=media"
	// if err := c.ShouldBindJSON(&card); err != nil {
	// 	utility.ValidationStruct(err, c)
	// 	return
	// }

	errFirebase := utility.UploadImageToFirebaseStorage(image)

	if errFirebase != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "errore in firebase"})
		return
	}

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

	if reflect.TypeOf(lat).String() != "string" || reflect.TypeOf(lat).String() != "string" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "latitude and longitude must be a string"})
		return
	}

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

func (repository *CardRepo) GetCardFromId(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var card models.Card
	err := models.GetCardFromId(repository.Db, &card, id)
	fmt.Println(err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Card not found"})
			return
		}
	}
	c.JSON(http.StatusOK, card)

}
