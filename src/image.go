package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"github.com/xiam/exif"
	"github.com/disintegration/imaging"
)

func createMiniature(dossier string,fileName string,pathMinDir string,size uint,response http.ResponseWriter, request *http.Request) {
	path := "static/galerie/" + dossier + "/" + fileName
	pathMinFile := pathMinDir + fileName
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
			fmt.Printf("new miniature dir!")
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
				print("resize : " + path + "\n")

				// decode jpeg into image.Image
				img, err := jpeg.Decode(file)
				if err != nil {
					log.Fatal("\n\roups:", err)
				}
				file.Close()

				// resize
				m := resize.Thumbnail(size, size, img, resize.Lanczos3)

				data, err := exif.Read(path)
				if err == nil {
					if(data.Tags["Orientation"] == "Bottom-right"){
						m = imaging.Rotate180(m)
					}
					if(data.Tags["Orientation"] == "Right-top"){
						m = imaging.Rotate270(m)
					}
					if(data.Tags["Orientation"] == "Left-bottom"){
						m = imaging.Rotate90(m)
					}
				}

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


func miniature(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	pathMinDir := "static/galerie/" + vars["dossier"] + "/min/"

	createMiniature(vars["dossier"],vars["file"],pathMinDir,300,response,request)

}


func bigMiniature(response http.ResponseWriter, request *http.Request) {
	
		vars := mux.Vars(request)
		pathMinDir := "static/galerie/" + vars["dossier"] + "/bigMin/"
	
		createMiniature(vars["dossier"],vars["file"],pathMinDir,900,response,request)
	
}
	