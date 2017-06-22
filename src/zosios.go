package main

import (
    "net/http"
	"github.com/gorilla/mux"
    "html/template"
    "log"
    "path/filepath"
    "io/ioutil"
)

type DetailGalerie struct {
    Titre    string
    Fichiers []string
}

var router = mux.NewRouter()

func main(){
    //liste des images pour le slider
    //files, _ := filepath.Glob("static/galerie/top");
    

    router.HandleFunc("/", index)
    router.HandleFunc("/index.html", index)
    router.HandleFunc("/galerie.html", galerie)
    router.HandleFunc("/couleur.html", couleur)
    router.HandleFunc("/galerie/{dossier}.html", detailGalerie)
    router.HandleFunc("/licenses.html", licenses)

    s1 := http.StripPrefix("/static/",http.FileServer(http.Dir("./static/")))
    router.PathPrefix("/").Handler(s1)
    http.Handle("/", router)
    http.ListenAndServe(":12345", nil)
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

func licenses(response http.ResponseWriter, request *http.Request) {
    
    //colors := []string{"#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"}
    colors := []string{"#fef9d1","#79a50a","#fd966d","#1a8ffb","#6b59cd","#38d9da","#ff3299","#e7bd27","fb02fe"}
    t := template.New("Label de ma template")
    
    t = template.Must(t.ParseFiles("pages/template.html", "pages/licenses.html"))
    err := t.ExecuteTemplate(response, "page", colors)
 
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}

func couleur(response http.ResponseWriter, request *http.Request) {
    
    //colors := []string{"#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"}
    colors := []string{"#fef9d1","#79a50a","#fd966d","#1a8ffb","#6b59cd","#38d9da","#ff3299","#e7bd27","fb02fe"}
    t := template.New("Label de ma template")
    
    t = template.Must(t.ParseFiles("pages/template.html", "pages/couleur.html"))
    err := t.ExecuteTemplate(response, "page", colors)
 
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}


func galerie(response http.ResponseWriter, request *http.Request) {
    files, _ := ioutil.ReadDir("./static/galerie")
    names_files := []string{}
    
    for _, f := range files {
            names_files = Extend(names_files,f.Name())
            
    }
    //colors := []string{"#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"}
    t := template.New("Label de ma template")
    
    t = template.Must(t.ParseFiles("pages/template.html", "pages/galerie.html"))
    err := t.ExecuteTemplate(response, "page", names_files)
 
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}

func detailGalerie(response http.ResponseWriter, request *http.Request){
    vars := mux.Vars(request)
    files, _ := filepath.Glob("static/galerie/"+vars["dossier"]+"/*");

    p := DetailGalerie{vars["dossier"], files}

    t := template.New("Label de ma template")
    t = template.Must(t.ParseFiles("pages/template.html", "pages/detailGalerie.html"))
    err := t.ExecuteTemplate(response, "page", p)
    
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

