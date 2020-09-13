package main // import "github.com/sneakstarberry/middle_w"

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sneakstarberry/middle_w/mw"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/").
		HandlerFunc(endpointHandler)

	n := negroni.New()
	n.Use(&mw.Logger{})
	n.UseHandler(router)

	err := http.ListenAndServe(":8080", n)
	if err != nil {
		panic(err)
	}
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint handler called")
}
