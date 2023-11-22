package controller

import (
	"context"
	"dego/config"
	"dego/person"
	"dego/pkg/client/postgresql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

func Enrich(c *gin.Context, p *person.Person) error {
	var err error
	links := config.NewLinks()
	resp := &http.Response{}
	ch := make(chan *http.Response, len(links))
	for _, v := range links {
		v := v
		go func() chan *http.Response {
			link := fmt.Sprintf("%s%s", v, p.Name)
			resp, err = http.Get(link)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return nil
			}
			ch <- resp
			return ch
		}()

	}
	mu := sync.Mutex{}
	for range links {
		mu.Lock()
		r := <-ch
		err = json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		}
		err = r.Body.Close()
		if err != nil {
			return err
		}
		mu.Unlock()
	}
	return nil
}

func Create(c *gin.Context) {
	client, err := postgresql.NewClient(context.TODO(), 3)
	if err != nil {
		log.Fatal(err)
	}
	repo := person.NewRepository(client)

	var p person.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	if err := Enrich(c, &p); err != nil {
		c.JSON(404, gin.H{"user not added": p})
	}

	id, err := repo.Create(context.TODO(), p)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"user not added": p})
		return
	}
	p.ID = id

	c.JSON(http.StatusOK, gin.H{"user added": p})
}
