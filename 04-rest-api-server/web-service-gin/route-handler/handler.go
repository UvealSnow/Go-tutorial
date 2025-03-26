package handler

import (
	"net/http"

	"example.com/db-interface/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repos *repository.RepositoryMap
}

func NewHandler(repos *repository.RepositoryMap) *Handler {
	handler := new(Handler)
	handler.repos = repos
	return handler
}

func (h *Handler) GetAlbums(c *gin.Context) {
	albums, err := h.repos.Albumrepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, albums)
}
