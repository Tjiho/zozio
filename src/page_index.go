package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(response http.ResponseWriter, request *http.Request) {

	names_files := []string{}

	data := struct {
		Files      []string
		Links      []Link
		Nav        bool
		Content_id string
	}{
		names_files,
		[]Link{},
		false,
		"index",
	}

	t := template.New("")
	t = template.Must(t.ParseFiles("pages/template.html", "pages/index.html", "pages/header-title.html"))
	err := t.ExecuteTemplate(response, "page", data)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}
