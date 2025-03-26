package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) GetAlbum(c *gin.Context) {
	id := c.Param("id")
	albumId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("GetAlbum: ID given is not numeric: %v", id)})
	}
	album, err := h.repos.Albumrepository.FindById(albumId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("GetAlbum: Unable to find album with ID: %v", albumId)})
	}
	c.JSON(http.StatusOK, album)
}
