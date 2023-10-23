package main

import (
	"dego/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/create", controller.Create)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}

}
