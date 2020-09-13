package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctr UserController) Signin(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user_id", 1)
	session.Set("user_email", "demo@demo.com")
	session.Set("user_username", "demo")
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "User signed in", "user": "demo"})
}

func (ctrl UserController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Signed out..."})
}
