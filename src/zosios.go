package main

import (
    "net/http"
    "net/url"
	"github.com/gorilla/mux"
    //"github.com/gorilla/securecookie"
    "html/template"
    "log"
    "io/ioutil"
)

type Link struct {
    Image   string
    Link    string
}
/*
var cookieHandler = securecookie.New(
     securecookie.GenerateRandomKey(64),
     securecookie.GenerateRandomKey(32)
)*/

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
    router.HandleFunc("/login.html", login)

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
    links := []Link{
        Link{
            Image:  "/static/images/home.svg",
            Link:   "/index.html",
        },
    }
    for _, f := range files {
            names_files = Extend(names_files,f.Name())
            
    }
    

    data := struct {
        Files []string
        Links []Link
        Nav bool
        Content_id string
    } {
        names_files,
        links,
        true,
        "flou",
    }
        
    
    //colors := []string{"#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"}
    t := template.New("Label de ma template")
    
    t = template.Must(t.ParseFiles("pages/template.html", "pages/galerie.html","pages/header-menu.html"))
    err := t.ExecuteTemplate(response, "page", data)
 
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}

func detailGalerie(response http.ResponseWriter, request *http.Request){
    vars := mux.Vars(request)
    path := "static/galerie/"+vars["dossier"];
    files, _ := ioutil.ReadDir(path)
    names_files := []string{}
    for _, f := range files {
            names_files = Extend(names_files,path+"/"+f.Name())            
    }


    links := []Link{
        Link{
            Image:  "/static/images/retour.svg",
            Link:   "/galerie.html",
        },
        Link{
            Image:  "/static/images/home.svg",
            Link:   "/index.html",
        },
        Link{
            Image:  "/static/images/upload.svg",
            Link:   "/index.html",
        },
        
    }

    data := struct {
        Files []string
        Title string
        Links []Link
        Nav bool
        Content_id string
    } {
        names_files,
        vars["dossier"],
        links,
        true,
        "flou",
    }

    t := template.New("Label de ma template")
    t = template.Must(t.ParseFiles("pages/template.html", "pages/detailGalerie.html","pages/header-menu.html"))
    err := t.ExecuteTemplate(response, "page", data)
    
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}

func index(response http.ResponseWriter, request *http.Request) {

    //files, _ := filepath.Glob("static/galerie/top/*");
    files, _ := ioutil.ReadDir("./static/galerie/top")
    names_files := []string{}
    for _, f := range files {
            names_files = Extend(names_files,"/static/galerie/top/"+f.Name())            
    }
    t := template.New("Label de ma template")

    data := struct {
        Files []string
        Links []Link
        Nav bool
        Content_id string
    } {
        names_files,
        []Link{},
        false,
        "index",
    }
    
    t = template.Must(t.ParseFiles("pages/template.html", "pages/index.html","pages/header-title.html"))
    err := t.ExecuteTemplate(response, "page", data)
 
    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}


func login(response http.ResponseWriter, request *http.Request) {
    //name := request.FormValue("login")
    //pass := request.FormValue("password")
    redirectTarget := "/"
   
        resp, err := http.PostForm( "http://localhost:8080/api/v1/users",
                                    url.Values{ "name": {"toto"}, 
                                                "surname": {"titi"},
                                                "password":{"123456789123"},
                                                "password_confirmation":{"1234567891"},
                                                "pseudo":{"toto"}})
        if err != nil {
	        log.Fatalf("Template execution: %s", err)
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        log.Fatalf("Template execution: %s", body)
        
        
        
        //setSession(name, response)
        redirectTarget = "/galerie"
    
    http.Redirect(response, request, redirectTarget, 302)
}
 
func logoutHandler(response http.ResponseWriter, request *http.Request) {
    //clearSession(response)
    http.Redirect(response, request, "/", 302)
}
