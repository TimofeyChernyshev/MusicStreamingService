package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Next()
	})

	r.GET("/stream", func(c *gin.Context) {
		streamHandler(c.Writer, c.Request)
	})

	log.Fatal(r.Run(":3001"))
}
