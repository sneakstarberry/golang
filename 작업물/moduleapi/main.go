package main // import "github.com/sneakstarberry/moduleapi"

import (
	log "github.com/sirupsen/logrus"
	"github.com/sneakstarberry/moduleapi/app"
)

func main() {
	log.WithFields(log.Fields{
		"Application": "TestApp",
		"Version:":    "v0.1.1",
	}).Info("Application start...")
	myapp := app.New("test app")
	myapp.Run()
}
