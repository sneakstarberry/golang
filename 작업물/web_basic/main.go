package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	s := NewServer()

	s.HandleFunc("GET", "/", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "welcome!")
	})

	s.HandleFunc("GET", "/", func(c *Context) {
		c.RenderTemplate("/public/index.html", map[string]interface{}{"time": time.Now()})
	})

	s.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
	})

	s.HandleFunc("GET", "/users/:id", func(c *Context) {
		u := User{Id: c.Params["id"].(string)}
		c.RenderXml(u)
	})

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
		c.RenderJson(u)
	})

	s.HandleFunc("GET", "/login", func(c *Context) {
		c.RenderTemplate("/public/login.html", map[string]interface{}{"message": "로그인이 필요합니다."})
	})

	s.HandleFunc("POST", "/login", func(c *Context) {
		if CheckLogin(c.Params["username"].(string), c.Params["password"].(string)) {
			http.SetCookie(c.ResponseWriter, &http.Cookie{
				Name:  "X_AUTH",
				Value: Sign(VerifyMessage),
				Path:  "/",
			})
			c.Redirect("/")
		}
		c.RenderTemplate("/public/login.html", map[string]interface{}{"message": "id또는 password가 일치하지 않습니다."})
	})

	s.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.Use(AuthHandler)

	s.Run(":8080")
}
