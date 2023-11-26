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

func Update(c *gin.Context) {
	client, err := postgresql.New(context.TODO(), 3)
	if err != nil {
		log.Fatal(err)
	}

	repo := person.NewRepository(client)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal("id input error")
	}
	var p person.Person
	p, err = repo.One(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"person not found with id": id})
		return
	}

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}

	err = repo.Update(context.TODO(), p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user not added": p})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user updated": p})
}
