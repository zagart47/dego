package controller

import (
	"context"
	"dego/person"
	"dego/pkg/client/postgresql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	client, err := postgresql.NewClient(context.TODO(), 3)
	repo := person.NewRepository(client)

	var p person.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	links := []string{
		"https://api.agify.io/?name=",
		"https://api.genderize.io/?name=",
	}

	for _, v := range links {
		link := fmt.Sprintf("%s%s", v, p.Name)
		resp, err := http.Get(link)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
			resp.Body.Close()
		}
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
			return
			resp.Body.Close()
		}

	}
	id, err := repo.Create(context.TODO(), p)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"user added with id": id})
}
