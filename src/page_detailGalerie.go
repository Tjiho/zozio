package main

import (
	"net/http"
	//"net/url"

	"github.com/gorilla/mux"
	//"github.com/gorilla/securecookie"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
)

func detailGalerie(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	path := "static/galerie/" + vars["dossier"]
	files, _ := ioutil.ReadDir(path)
	names_files := []string{}

	//list files
	for _, f := range files {
		var extension = filepath.Ext(f.Name())
		if !f.IsDir() && (extension == ".jpg" || extension == ".JPG") {
			names_files = Extend(names_files, vars["dossier"]+"/"+f.Name())
		}

	}

	//create menu
	links := []Link{
		Link{
			Image: "/static/images/retour.svg",
			Link:  "/galerie.html",
		},
		Link{
			Image: "/static/images/home.svg",
			Link:  "/index.html",
		},
		Link{
			Image: "/static/images/upload.svg",
			Link:  "/index.html",
		},
	}

	data := struct {
		Files      []string
		Title      string
		Links      []Link
		Nav        bool
		Content_id string
	}{
		names_files,
		vars["dossier"],
		links,
		true,
		"photos",
	}

	t := template.New("")
	t = template.Must(t.ParseFiles("pages/template.html", "pages/detailGalerie.html", "pages/header-menu.html"))
	err := t.ExecuteTemplate(response, "page", data)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}