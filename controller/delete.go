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

func Delete(c *gin.Context) {
	client, err := postgresql.NewClient(context.TODO(), 3)
	repo := person.NewRepository(client)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal("id input error")
	}

	err = repo.Delete(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"person not found with id": id})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted person with id": id})
}
