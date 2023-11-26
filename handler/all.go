package handler

import (
	"context"
	"dego/person"
	"dego/pkg/client/postgresql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func All(c *gin.Context) {
	client, err := postgresql.New(context.TODO(), 3)
	if err != nil {
		log.Fatal(err)
	}

	repo := person.NewRepository(client)

	id, err := repo.All(context.TODO())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"all persons": id})
}
