package main //import "github.com/sneakstarberry/web2"

import (
	"net/http"

	"github.com/gorilla/mux"
	postapi "github.com/sneakstarberry/web2/api/post_api"

	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/api/posts/", postapi.FindAll).Methods("GET")

	router.HandleFunc("/api/posts/", postapi.Create).Methods("POST")

	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("/static/")))
	n.UseHandler(router)

	n.Run(":8000")
}
