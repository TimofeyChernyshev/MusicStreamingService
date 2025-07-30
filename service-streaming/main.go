package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Next()
	})

	r.GET("/stream", func(c *gin.Context) {
		streamHandler(c.Writer, c.Request)
	})

	log.Fatal(r.Run(":3001"))
}
