package app

import (
	"dego/config"
	"dego/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/create", controller.Create)
	router.POST("/update/:id", controller.Update)
	router.GET("/all", controller.GetAll)
	router.GET("/one/:id", controller.GetOne)
	router.GET("/delete/:id", controller.Delete)
	server := config.NewConfig()
	dsn := fmt.Sprintf("%s:%s", server.HTTPHost, server.HTTPPort)
	err := router.Run(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

}
