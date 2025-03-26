package main

import (
	"log"
	"net/http"

	dbinterface "example.com/db-interface"
	"example.com/db-interface/repository"
	"github.com/gin-gonic/gin"
)

type handler struct {
	repos *repository.RepositoryMap
}

func newHandler(repos *repository.RepositoryMap) *handler {
	handler := new(handler)
	handler.repos = repos
	return handler
}

func (h *handler) getAlbums(c *gin.Context) {
	albums, err := h.repos.Albumrepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, albums)
}

func main() {
	repositories, err := dbinterface.ConnectToDB()
	handler := newHandler(repositories)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	router := gin.Default()
	router.GET("/albums", handler.getAlbums)

	router.Run("localhost:3000")
}
