package middleware

import (
	"Bear997/api/auth"
	"Bear997/api/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	userRepo := controllers.User()
	token := c.GetHeader("Authorization")
	fmt.Println(token)
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	auth.ValidateJwt(token, c, userRepo.Db)
	c.Next()

}
