package main

import (
	"net/http"

	"github.com/sneakstarberry/web3/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
