package db

import (
	"log"
	"net/http"

	"github.com/TimofeyChernyshev/MusicStreamingService/models"
	"github.com/gin-gonic/gin"
)

func GetTracks(c *gin.Context) {
	var songs []models.Song

	if err := models.AppDB.Find(&songs).Error; err != nil {
		log.Println("Database query error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func GetAlbums(c *gin.Context) {
	var albums []models.Album

	if err := models.AppDB.Find(&albums).Error; err != nil {
		log.Println("Database query error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}
