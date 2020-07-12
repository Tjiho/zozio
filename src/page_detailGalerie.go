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
	"strings"
	"path/filepath"
	"github.com/xiam/exif"
)


func getExifDate(file ImageSimple) string{
    value, ok := file.ExifData["Date and Time"]
	if ok {
		return value
    } else {
		return "2099:12:30 23:50:50"
    }
}

func sortImagesByDateInsertion(files []ImageSimple,len int) []ImageSimple{
	j := 0
	var y ImageSimple
	var x string

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
			data1, err1 := exif.Read("galerie/"+files[j])
			data2, err2 := exif.Read("galerie/"+files[j+1])
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

func passwdPage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	data := struct {
		Title      string
		Nav        bool
		Content_id string
		NightMode bool
		Dossier string
	}{
		"Page protégée",
		true,
		"passwd",
		getNightValue(request),
		vars["dossier"],
	}

	t := template.New("")

	t = template.Must(t.ParseFiles("pages/template.html", "pages/password.html", "pages/header-menu.html"))
	err := t.ExecuteTemplate(response, "page", data)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func detailGalerie(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	session, session_err := store.Get(request, "zozio")
	if session_err != nil {
			http.Error(response, session_err.Error(), http.StatusInternalServerError)
			return
	}
	path := "galerie/" + vars["dossier"]
	path_passwd_file := path + "/passwd.txt"
	authorized := true
	data, err := ioutil.ReadFile(path_passwd_file)
	trimData := strings.TrimSpace(string(data))

  if err == nil {
			authorized = false
			session_auth, session_done := session.Values["access-" + vars["dossier"]].(bool)
			if (session_auth && session_done)	{
				authorized = true
			} else {
				err = request.ParseForm()
				if err != nil {
					response.WriteHeader(400)
			    return
				}
				password := request.Form.Get("password")
				if (string(password) == string(trimData)) {
					authorized = true
					session.Values["access-" + vars["dossier"]] = true
			    session.Save(request, response)
				}
			}
  }

  if !authorized {
		response.WriteHeader(401)
		passwdPage(response,request)
	}

	if authorized	{
		files, _ := ioutil.ReadDir(path)
		names_files := []ImageSimple{}
		i := 0
		//list files
		for _, f := range files {
			var extension = filepath.Ext(f.Name())

			if !f.IsDir() && (extension == ".jpg" || extension == ".JPG" || extension == ".png" || extension == ".PNG") {
				data, err := exif.Read(path+ "/" + f.Name())
				if err == nil {
					names_files = append(names_files, ImageSimple { //Todo: use a list and not a Slice
						Name: f.Name(),
						Path: vars["dossier"]+"/"+f.Name(),
						ExifData: data.Tags,
					})
				} else {
					names_files = append(names_files, ImageSimple { //Todo: use a list and not a Slice
						Name: f.Name(),
						Path: vars["dossier"]+"/"+f.Name(),
						ExifData: nil,
					})
				}
				i = i+1

			}

		}
		names_files = sortImagesByDateInsertion(names_files,i)

		//create menu
		links := []Link{
			Link{
				Image: "/static/images/home.svg",
				Link:  "/",
			},
		}

		data := struct {
			Files      []ImageSimple
			Title      string
			Links      []Link
			Nav        bool
			Content_id string
			NightMode bool
		}{
			names_files,
			vars["dossier"],
			links,
			true,
			"photos",
			getNightValue(request),
		}

		t := template.New("")
		t = template.Must(t.ParseFiles("pages/template.html", "pages/detailGalerie.html", "pages/header-menu.html"))
		err := t.ExecuteTemplate(response, "page", data)

		if err != nil {
			log.Fatalf("Template execution: %s", err)
		}
	}
}
