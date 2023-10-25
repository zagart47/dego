package controller

import (
	"context"
	"dego/person"
	"dego/pkg/client/postgresql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOne(c *gin.Context) {
	client, err := postgresql.NewClient(context.TODO(), 3)
	repo := person.NewRepository(client)

	id := c.Param("id")

	oneId, err := repo.FindOne(context.TODO(), id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"one person": oneId})
}
