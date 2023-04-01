package main

import (
	"Bear997/api/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf(os.Getenv("MYSQL_DATABASE"))
	r := gin.Default()
	r.GET("api/v1/dio", response)
	cardRepo := controllers.Card()
	r.POST("api/v1/createCard", cardRepo.CreateCard)
	r.GET("api/v1/card/locate", cardRepo.GetCardFromPosition)
	r.GET("api/v1/card/:id", cardRepo.GetCardFromId)
	r.Run(":3000")
}
func response(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "tascioeteroh"})
}
