package main

import (
	"net/http"
)

func nightMode(response http.ResponseWriter, request *http.Request) {
    session, err := store.Get(request, "zozio")
    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }
    err = request.ParseForm()
	if err != nil {
		response.WriteHeader(400)
        return
	}

	nightMode := request.Form.Get("nightMode")

    session.Values["nightMode"] = (nightMode == "1")
    session.Save(request, response)
    response.WriteHeader(201)
}

func getNightValue(request *http.Request) bool{
    session, _ := store.Get(request, "zozio")
    nightMode, ok := session.Values["nightMode"].(bool)
    return (ok && nightMode)
}
