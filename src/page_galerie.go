package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func galerie(response http.ResponseWriter, request *http.Request) {

	// session, err := store.Get(request, "zozio")
    // if err != nil {
    //    http.Error(response, err.Error(), http.StatusInternalServerError)
    //    return
    //}

    // isConnected, _ := session.Values["connected"].(bool)


	files, _ := ioutil.ReadDir("./galerie")
	names_files := []string{}
	nightMode := getNightValue(request)

	for _, f := range files {
		if _, err := os.Stat("./galerie/"+f.Name()+"/private.txt");  os.IsNotExist(err) {
			names_files = Extend(names_files, f.Name())
		}

	}

	data := struct {
		Files      []string
		Title      string
		Nav        bool
		Content_id string
		NightMode bool
	}{
		names_files,
		"Les albums photo",
		true,
		"galeries",
		nightMode,
	}

	//colors := []string{"#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"}
	t := template.New("")

	t = template.Must(t.ParseFiles("pages/template.html", "pages/galerie.html", "pages/header-menu.html"))
	err := t.ExecuteTemplate(response, "page", data)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}
