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
	fmt.Printf("\ndslkvbnaksdljbv")
	r := gin.Default()
	r.GET("/dio", response)
	cardRepo := controllers.Card()
	r.POST("api/v1/createCard", cardRepo.CreateCard)
	r.Run(":3000")
}
func response(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "tasciogay"})
}
