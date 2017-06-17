package main

import (
    "net/http"
	"github.com/gorilla/mux"
    "html/template"
    "log"
    "path/filepath"
)
type Page struct {
    Title    string
    Articles []string
}
var router = mux.NewRouter()

func main(){
    //liste des images pour le slider
    //files, _ := filepath.Glob("static/galerie/top");
    

    router.HandleFunc("/", index)
    router.HandleFunc("/galerie", galerie)
    s1 := http.StripPrefix("/static/",http.FileServer(http.Dir("./static/")))
    router.PathPrefix("/").Handler(s1)
    http.Handle("/", router)
    http.ListenAndServe(":12345", nil)
}

func galerie(response http.ResponseWriter, request *http.Request) {
    files, _ := filepath.Glob("static/galerie/top/*");
    t := template.New("Label de ma template")
    
    t = template.Must(t.ParseFiles("pages/template.html", "pages/galerie.html"))
    err := t.ExecuteTemplate(response, "page", files)
 
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}


func index(response http.ResponseWriter, request *http.Request) {
    files, _ := filepath.Glob("static/galerie/top/*");
    t := template.New("Label de ma template")
    
    t = template.Must(t.ParseFiles("pages/template.html", "pages/index.html"))
    err := t.ExecuteTemplate(response, "page", files)
 
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}

