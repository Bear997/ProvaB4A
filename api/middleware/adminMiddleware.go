package middleware

import (
	"Bear997/api/auth"
	"Bear997/api/controllers"
	"Bear997/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(c *gin.Context) {
	var user models.User
	userRepo := controllers.User()
	token := c.GetHeader("Authorization")
	id := auth.GetIdFromToken(token)
	if id == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizeed"})
		return
	}
	err := models.GetUserFromId(userRepo.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizeed"})
		return
	}
	if user.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizeed"})
		return
	}
	c.Next()
}
