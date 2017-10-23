package main

import (
	"net/http"
)


func login(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("login")
	//pass := request.FormValue("password")
	/*
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

	//setSession(name, response)*/


	redirectTarget := "/galerie.html"
	
	session, err := store.Get(request, "zozio")
    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }

    session.Values["login"] = name
    session.Values["connected"] = true

    session.Save(request, response)

	http.Redirect(response, request, redirectTarget, 302)
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	//clearSession(response)
	http.Redirect(response, request, "/", 302)
}
