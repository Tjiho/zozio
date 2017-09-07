package main

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	//"github.com/gorilla/securecookie"
	"fmt"
	"html/template"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

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

	s1 := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
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

func galerie(response http.ResponseWriter, request *http.Request) {

	files, _ := ioutil.ReadDir("./static/galerie")
	names_files := []string{}

	links := []Link{
		Link{
			Image: "/static/images/home.svg",
			Link:  "/index.html",
		},
	}

	for _, f := range files {
		names_files = Extend(names_files, f.Name())
	}

	data := struct {
		Files      []string
		Links      []Link
		Nav        bool
		Content_id string
	}{
		names_files,
		links,
		true,
		"galeries",
	}

	//colors := []string{"#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"}
	t := template.New("")

	t = template.Must(t.ParseFiles("pages/template.html", "pages/galerie.html", "pages/header-menu.html"))
	err := t.ExecuteTemplate(response, "page", data)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

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

/*********** Acceuil **********/
/******************************/
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

func categorie(response http.ResponseWriter, request *http.Request) {

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
		"categories",
	}

	t := template.New("Label de ma template")
	t = template.Must(t.ParseFiles("pages/template.html", "pages/categorie.html", "pages/header-title.html"))
	err := t.ExecuteTemplate(response, "page", data)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func login(response http.ResponseWriter, request *http.Request) {
	//name := request.FormValue("login")
	//pass := request.FormValue("password")
	redirectTarget := "/"

	resp, err := http.PostForm("http://localhost:8080/api/v1/users",
		url.Values{"name": {"toto"},
			"surname":               {"titi"},
			"password":              {"123456789123"},
			"password_confirmation": {"1234567891"},
			"pseudo":                {"toto"}})
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
