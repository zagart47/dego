package handler

import (
	"context"
	"dego/person"
	"dego/pkg/client/postgresql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func One(c *gin.Context) {
	client, err := postgresql.New(context.TODO(), 3)
	if err != nil {
		log.Fatal(err)
	}
	repo := person.NewRepository(client)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal("id input error")
	}

	p, err := repo.One(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"person not found with id": id})
	}
	if p.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"one person": p})
	}
}
