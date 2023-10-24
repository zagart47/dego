package controller

import (
	"context"
	"dego/person"
	"dego/pkg/client/postgresql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	client, err := postgresql.NewClient(context.TODO(), 3)
	repo := person.NewRepository(client)

	id, err := repo.FindAll(context.TODO())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"all persons": id})
}
