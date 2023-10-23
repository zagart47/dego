package main

import (
	"dego/controller"
	"dego/router"
	"log"
)

func main() {
	router.Router.POST("/create", controller.Create)
	err := router.Router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
