package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	router *mux.Router
}

func New(name string) *Application {
	app := &Application{}
	app.router = mux.NewRouter()
	return app
}

func (a *Application) Run() {
	a.router.HandleFunc("/user/hello", Handler_Hello)
	http.Handle("/", a.router)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Handler_Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
