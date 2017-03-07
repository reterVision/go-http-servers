package gin

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// Serve starts a HTTP server written with gin
func Serve(addr string) {
	r := gin.Default()
	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Print(r.Run(addr))
}
