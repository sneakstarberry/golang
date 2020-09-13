package main // import "github.com/sneakstarberry/file"

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	fileType := strings.Split(handler.Filename, ".")

	tempFile, err := ioutil.TempFile("temp-images", "upload-*."+fileType[len(fileType)-1])
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	fmt.Fprintf(w, "print: %s", tempFile.Name())
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8000", nil)
}

func main() {
	fmt.Println("hello World")
	setupRoutes()
}
