package main

import (
	"log"

	dbinterface "example.com/db-interface"
	handler "example.com/web-service-gin/route-handler"
	"github.com/gin-gonic/gin"
)

func main() {
	repositories, err := dbinterface.ConnectToDB()
	handler := handler.NewHandler(repositories)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbum)
	router.POST("/albums", handler.CreateAlbum)

	router.Run("localhost:3000")
}
