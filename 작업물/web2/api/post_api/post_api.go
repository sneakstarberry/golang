package postapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sneakstarberry/web2/config"
	"github.com/sneakstarberry/web2/entities"
	"github.com/sneakstarberry/web2/model"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		postModel := model.PostModel{
			Db: db,
		}
		posts, err2 := postModel.FindAll(request)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, posts)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	file, handler, err := request.FormFile("image")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fileType := strings.Split(handler.Filename, ".")

	tempFile, err := ioutil.TempFile("static", "upload-*."+fileType[len(fileType)-1])
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	var post entities.Post

	post.Image = tempFile.Name()
	post.Title = request.FormValue("title")
	post.Content = request.FormValue("content")
	db, err2 := config.GetDB()
	if err2 != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		postModel := model.PostModel{
			Db: db,
		}
		err3 := postModel.Create(&post)
		if err3 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, post)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
