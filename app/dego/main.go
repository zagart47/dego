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
	router.GET("/all", controller.GetAll)
	router.GET("/one/:id", controller.GetOne)
	err := router.Run("localhost:8081")
	if err != nil {
		log.Fatal(err.Error())
	}

}
