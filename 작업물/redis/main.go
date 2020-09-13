package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		if sessionID == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "not authed",
			})
			c.Abort()
		}
	}
}

var (
	//RedisHost ...
	RedisHost = os.Getenv("REDISHOST")
	//RedisPort ...
	RedisPort = os.Getenv("REDISPORT")
)

func main() {
	r := gin.Default()

	store, _ := redis.NewStore(10, "tcp", RedisHost+":"+RedisPort, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	userController := new(UserController)

	r.POST("/signin", userController.Signin)
	r.GET("/signout", userController.Signout)

	auth := r.Group("/auth")
	auth.Use(AuthRequired())
	{
		auth.GET("/ping", AuthRequired(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
				"host":    RedisHost,
			})
		})
	}

	r.Run(":8080")
}
