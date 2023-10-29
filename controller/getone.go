package controller

import (
	"context"
	"dego/person"
	"dego/pkg/client/postgresql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetOne(c *gin.Context) {
	client, err := postgresql.NewClient(context.TODO(), 3)
	if err != nil {
		log.Fatal(err)
	}
	repo := person.NewRepository(client)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal("id input error")
	}

	person, err := repo.FindOne(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"person not found with id": id})
	}
	if person.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"one person": person})
	}
}
