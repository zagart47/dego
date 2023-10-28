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
	router.POST("/update/:id", controller.Update)
	router.GET("/all", controller.GetAll)
	router.GET("/one/:id", controller.GetOne)
	router.GET("/delete/:id", controller.Delete)
	err := router.Run("localhost:8081")
	if err != nil {
		log.Fatal(err.Error())
	}

}
