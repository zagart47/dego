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
	repo := person.NewRepository(client)

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal("error with id input")
	}

	oneId, err := repo.FindOne(context.TODO(), int64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"one person": oneId})
}
