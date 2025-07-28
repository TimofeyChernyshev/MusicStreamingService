package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/TimofeyChernyshev/MusicStreamingService/db"
	"github.com/TimofeyChernyshev/MusicStreamingService/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.StaticFile("/", "./frontend/player/index.html")
	r.Static("/player-static", "./frontend/player")
	r.Static("/storage/cover", "./storage/cover")

	config := models.Config{
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		Admin_User:     os.Getenv("DB_USER_ADMIN"),
		Admin_Password: os.Getenv("DB_PASSWORD_ADMIN"),
		DBName:         os.Getenv("DB_NAME"),
		SSLMode:        os.Getenv("DB_SSLMODE"),
		App_User:       os.Getenv("DB_USER"),
		App_Password:   os.Getenv("DB_PASSWORD"),
	}

	models.InitDB(config)

	r.GET("api/albums", func(c *gin.Context) {
		db.GetAlbums(c)
	})

	r.GET("api/albums/:id/tracks", func(c *gin.Context) {
		db.GetTracks(c)
	})

	r.GET("api/stream", func(c *gin.Context) {
		track := c.Query("track")
		if track == "" {
			log.Println("error: track parameter is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "track parameter is required"})
			return
		}

		streamingURL := fmt.Sprintf(os.Getenv("STREAM_URL")+"/stream?track=%s", url.QueryEscape(track))

		resp, err := http.Get(streamingURL)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			c.Header(k, v[0])
		}
		c.Status(resp.StatusCode)
		c.Stream(func(w io.Writer) bool {
			io.Copy(w, resp.Body)
			return false
		})
	})

	log.Fatal(r.Run(":8080"))
}
