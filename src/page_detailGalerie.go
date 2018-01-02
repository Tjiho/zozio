package main

import (
	"net/http"
	"fmt"
	//"net/url"
	//"time"
	"github.com/gorilla/mux"
	//"github.com/gorilla/securecookie"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"github.com/xiam/exif"
)


func getExifDate(file string) string{
	data, err := exif.Read("static/galerie/"+file)
    if err != nil {
        return "2099:12:30 23:50:50"
    } else {
        return data.Tags["Date and Time (Original)"]
    }
}

func sortImagesByDateInsertion(files []string,len int) []string{
	j := 0
	x := ""
	y := ""
	for i:= 1; i < len - 1;i++ {
	    x = getExifDate(files[i])
		y = files[i]
		j = i
        for j > 0 &&  getExifDate(files[j-1]) > x {
            files[j] = files[j - 1]
            j = j-1
        }
		files[j] = y 
	}
	return files
}

func sortImagesByDate(files []string,len int) []string{
	for i:= len - 1; i > 0;i-- {
		for j:= 0;j < i;j++ {
			data1, err1 := exif.Read("static/galerie/"+files[j])
			data2, err2 := exif.Read("static/galerie/"+files[j+1])
			if err1 == nil {
				if err2 == nil {
					//date1 = 		
					if(data2.Tags["Date and Time (Original)"] < data1.Tags["Date and Time (Original)"]) {
						//fmt.Printf("\n")
						//fmt.Printf(data2.Tags["Date and Time (Original)"] +" < ")
						//fmt.Printf(data1.Tags["Date and Time (Original)"])
						//fmt.Printf(" : ok")
						a := files[j]
						files[j] = files[j+1]
						files[j+1] = a
					}
				} else {
					fmt.Printf("\nno date - "+files[j+1])
				}
			} else {
				a := files[j]
				files[j] = files[j+1]
				files[j+1] = a
			}
		}
	}
	return files
}


func detailGalerie(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	path := "static/galerie/" + vars["dossier"]
	files, _ := ioutil.ReadDir(path)
	names_files := []string{}
	i := 0
	//list files
	for _, f := range files {
		var extension = filepath.Ext(f.Name())

		if !f.IsDir() && (extension == ".jpg" || extension == ".JPG") {
			names_files = Extend(names_files, vars["dossier"]+"/"+f.Name())
			i = i+1
		}

	}
	names_files = sortImagesByDateInsertion(names_files,i)

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
