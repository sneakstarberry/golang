package main

import (
	"net/http"

	"github.com/sneakstarberry/web1/myapp"
)

func main() {

	http.ListenAndServe(":8000", myapp.NewHttpHandler())
}
