package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// form3, _ := c.FormFile("file")
		form4, _ := c.MultipartForm()

		var form2 LoginForm
		c.ShouldBind(&form2)

		// or you can simply use autobinding with ShouldBind method:
		var form LoginForm
		// in this case proper binding will be automatically selected
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{
					"status":   "you are logged in",
					"user":     form2.User,
					"password": form2.Password,
					"test":     form4,
					"user2":    form4.Value["user"][0],
				})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run("0.0.0.0:8000")
}
