package app

import (
	"dego/config"
	"dego/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/create", handler.Create)
	router.PUT("/update/:id", handler.Update)
	router.GET("/all", handler.All)
	router.GET("/one/:id", handler.One)
	router.DELETE("/delete/:id", handler.Delete)
	server := config.NewServerConfig()
	dsn := fmt.Sprintf("%s:%s", server.HTTPHost, server.HTTPPort)
	err := router.Run(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

}
