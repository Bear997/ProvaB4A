package router

import (
	"Bear997/api/controllers"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(r *gin.Engine) {
	cardRepo := controllers.Card()
	r.POST("api/v1/createCard", cardRepo.CreateCard)
	r.GET("api/v1/card/locate", cardRepo.GetCardFromPosition)
	r.GET("api/v1/card/:id", cardRepo.GetCardFromId)
	r.Run(":3000")
}
