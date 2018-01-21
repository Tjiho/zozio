package main

import (
	"math/rand"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"path/filepath"
	//"github.com/xiam/exif"
)

func randomImage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	path := "static/galerie/" + vars["dossier"]
	files, _ := ioutil.ReadDir(path)
	names_files := []string{}
	i := 0
	//list files
	for _, f := range files {
		var extension = filepath.Ext(f.Name())

		if !f.IsDir() && (extension == ".jpg" || extension == ".JPG") {
			names_files = Extend(names_files, vars["dossier"]+"/"+f.Name())
			i = i+1
		}

	}

	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	randomInt := source.Intn(i)

	http.Redirect(response, request, "/static/galerie/"+names_files[randomInt], http.StatusTemporaryRedirect)
}