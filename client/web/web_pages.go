package web

import (
	"Battleships/client/data"
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Token string
}

func battlePageHandler(w http.ResponseWriter, r *http.Request) {
	var fileName = "assets/battle_page.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Token: "Game Id: " + data.GetToken(),
	}

	err = t.ExecuteTemplate(w, "battle_page", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/battle":
		battlePageHandler(w, r)
	default:

		fmt.Fprint(w, "Default")

	}
}
