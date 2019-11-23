package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
				print("go resize : " + path + "\n")
				extension := filepath.Ext(fileName)
				// decode jpeg or png into image.Image
				var img image.Image

				if  extension == ".jpg" || extension == ".JPG" {
					img, err = jpeg.Decode(file)
				} else if extension == ".png" || extension == ".PNG" {
					img, err = png.Decode(file)
				}

				if err != nil {
					log.Fatal("\n\roups:", err)
				}
				file.Close()

				// resize
				m := resize.Thumbnail(size, size, img, resize.Lanczos3)

				data, err := exif.Read(path)
				if err == nil {
					fmt.Println(data.Tags["Orientation"])
					if(data.Tags["Orientation"] == "Bottom-right"){
						m = imaging.Rotate180(m)
					}
					if(data.Tags["Orientation"] == "Right-top"){
						m = imaging.Rotate270(m)
					}
					if(data.Tags["Orientation"] == "Left-bottom"){
						m = imaging.Rotate90(m)
					}



					fmt.Println()
				}

				//create new image
				out, err := os.Create(pathMinFile)
				if err != nil {
					log.Fatal(err)
				}
				defer out.Close()
				if  extension == ".jpg" || extension == ".JPG" {
					jpeg.Encode(out, m, nil)
				} else if extension == ".png" || extension == ".PNG" {
					png.Encode(out, m)
				}
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

		createMiniature(vars["dossier"],vars["file"],pathMinDir,1100,response,request)
}
