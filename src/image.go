package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

func miniature(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	pathMinDir := "static/galerie/" + vars["dossier"] + "/min/"
	path := "static/galerie/" + vars["dossier"] + "/" + vars["file"]
	pathMinFile := pathMinDir + vars["file"]
	ok := true

	//get image

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("\n\roups:", err)
		response.WriteHeader(http.StatusNotFound)

		ok = false
	}

	//check if miniature dir exist
	stat, err := os.Stat(pathMinDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("new dir!")
			os.Mkdir(pathMinDir, os.FileMode(0755))
		} else {
			fmt.Printf("error!")
			response.WriteHeader(http.StatusInternalServerError)
			ok = false
		}
	} else {
		if !stat.IsDir() {
			fmt.Printf("error!")
			response.WriteHeader(http.StatusInternalServerError)
			ok = false
		}
	}

	if ok == true {

		//check if miniature file already exist
		_, err := os.Stat(pathMinFile)
		if err != nil {
			if os.IsNotExist(err) {
				print("go resize : " + path)

				// decode jpeg into image.Image
				img, err := jpeg.Decode(file)
				if err != nil {
					log.Fatal("\n\roups:", err)
				}
				file.Close()

				// resize
				m := resize.Thumbnail(300, 300, img, resize.Lanczos3)

				//create new image
				out, err := os.Create(pathMinFile)
				if err != nil {
					log.Fatal(err)
				}
				defer out.Close()
				jpeg.Encode(out, m, nil)
			} else {
				response.WriteHeader(http.StatusInternalServerError)
			}
		}
		http.Redirect(response, request, "/"+pathMinFile, http.StatusTemporaryRedirect)
	}

}
