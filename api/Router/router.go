package router

import (
	"Bear997/api/controllers"
	"Bear997/api/middleware"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(r *gin.Engine) {
	cardRepo := controllers.Card()
	userRepo := controllers.User()
	r.POST("api/v1/createCard", cardRepo.CreateCard)
	r.GET("api/v1/card/locate", cardRepo.GetCardFromPosition)
	r.GET("api/v1/card/:id", middleware.AuthMiddleware, cardRepo.GetCardFromId)
	r.POST("api/v1/signup", userRepo.CreateUser)
	r.POST("api/v1/login", userRepo.Login)

	r.Run(":3000")
}
