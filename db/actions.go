package db

import (
	"log"
	"net/http"

	"github.com/TimofeyChernyshev/MusicStreamingService/models"
	"github.com/gin-gonic/gin"
)

func (r *SongRepository) GetTracks(c *gin.Context) {
	albumID := c.Param("id")
	rows, err := r.db.Query("SELECT id, title, file_name FROM tracks where album_id = $1", albumID)
	if err != nil {
		log.Println("Database query error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tracks []models.Song
	for rows.Next() {
		var t models.Song
		err := rows.Scan(&t.ID, &t.TITLE, &t.FILE_NAME)
		if err != nil {
			log.Println("Row scanning error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tracks = append(tracks, t)
	}
	c.JSON(http.StatusOK, tracks)
}

func (r *AlbumRepository) GetAlbums(c *gin.Context) {
	rows, err := r.db.Query("SELECT id, title, artist, cover FROM albums")
	if err != nil {
		log.Println("Database query error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var a models.Album
		err := rows.Scan(&a.ID, &a.TITLE, &a.ARTIST, &a.COVER)
		if err != nil {
			log.Println("Row scanning error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		albums = append(albums, a)
	}
	c.JSON(http.StatusOK, albums)
}
