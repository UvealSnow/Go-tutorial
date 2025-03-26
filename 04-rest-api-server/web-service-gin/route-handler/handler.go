package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/db-interface/album"
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
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (h *Handler) GetAlbum(c *gin.Context) {
	id := c.Param("id")
	albumId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ID given is not numeric: %v", id)})
		return
	}
	album, err := h.repos.Albumrepository.FindById(albumId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Unable to find album with ID: %v", albumId)})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *Handler) CreateAlbum(c *gin.Context) {
	var newAlbum album.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	if newAlbum.Title == "" || newAlbum.Artist == "" || newAlbum.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields must be non-empty and price must be greater than 0"})
		return
	}

	albumId, err := h.repos.Albumrepository.Insert(&newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert", "data": newAlbum})
		return
	}

	album, err := h.repos.Albumrepository.FindById(albumId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Unable to find album with ID: %v", albumId)})
		return
	}
	c.JSON(http.StatusCreated, album)
}
