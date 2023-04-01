package main

import (
	router "Bear997/api/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.DefaultRouter(r)
}
