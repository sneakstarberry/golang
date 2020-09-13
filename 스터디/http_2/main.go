package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping/:id", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "count"+c.Param("id"),
		})
	})

	r.POST("/ping", func(c *gin.Context){
		
		c.String(200, c.Request.Body
	})
	r.PUT("/ping", func(c *gin.Context){
		
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/ping", func(c *gin.Context){
		
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	r.Run()
}
