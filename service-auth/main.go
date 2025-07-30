package main

import (
	"log"
	"os"

	"github.com/TimofeyChernyshev/MusicStreamingService/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		c.Writer.Header().Set("Access-Control-Allow-Methods", os.Getenv("ALLOWED_METHODS"))
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

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

	r.POST("/login", Login)
	r.POST("/signup", Signup)
	r.GET("/logout", Logout)

	log.Fatal(r.Run(":3002"))
}
