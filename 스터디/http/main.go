package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ping/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong" + c.Param("id"),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
