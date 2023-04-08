package router

import (
	"Bear997/api/controllers"
	"Bear997/api/middleware"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(r *gin.Engine) {
	cardRepo := controllers.Card()
	userRepo := controllers.User()

	//route for guest user
	guest := r.Group("api/v1")
	guest.POST("/signup", userRepo.CreateUser)
	guest.POST("/login", userRepo.Login)

	//route for logged user
	logged := guest.Group("/", middleware.AuthMiddleware)
	logged.GET("card/:id", cardRepo.GetCardFromId)
	logged.GET("card/locate", cardRepo.GetCardFromPosition)
	logged.GET("/allCards", userRepo.GetAllCardsOfUser)
	logged.GET("/logout")       //TODO to implement
	logged.PATCH("/updateUser") //TODO to implement change profile obviously from jwt token forse (nonloso)
	logged.POST("/assign/:cardId", cardRepo.AssignCardToUser)
	//route for admin user
	admin := logged.Group("/admin", middleware.AdminMiddleware)
	admin.POST("/createCard", cardRepo.CreateCard)
	admin.GET("/cards")            //TODO to implement get all cards
	admin.PATCH("/updateCard/:id") //TODO to implement update a card
	admin.DELETE("deleteCard/:id") //TODO to implement delete a card

	r.Run(":3000")
}
