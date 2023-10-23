package controller

import (
	"dego/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	links := []string{
		"https://api.agify.io/?name=",
		"https://api.genderize.io/?name=",
		"https://api.nationalize.io/?name=",
	}

	for _, v := range links {
		link := fmt.Sprintf("%s%s", v, person.Name)
		resp, err := http.Get(link)
		defer resp.Body.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = json.NewDecoder(resp.Body).Decode(&person)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"ready": person})
}
