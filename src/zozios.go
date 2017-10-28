package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("er378uei74783es75fiusfieh5i!sfeij*/$dq"))

type Link struct {
	Image string
	Link  string
}

var router = mux.NewRouter()

func main() {
	//liste des images pour le slider
	//files, _ := filepath.Glob("static/galerie/top");

	router.HandleFunc("/", index)
	router.HandleFunc("/index.html", index)
	router.HandleFunc("/galerie.html", galerie)
	router.HandleFunc("/galerie/{dossier}.html", detailGalerie)
	router.HandleFunc("/login.html", login)
	router.HandleFunc("/miniature/{dossier}/{file}", miniature)
	router.HandleFunc("/bigMiniature/{dossier}/{file}", bigMiniature)
	s1 := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/").Handler(s1)
	http.Handle("/", router)
	http.ListenAndServe("127.0.0.1:8764", nil)
}

func Extend(slice []string, element string) []string {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]string, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
