package web

import (
	"fmt"
	"html/template"
	"net/http"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func battle(w http.ResponseWriter, r *http.Request) {
	var fileName = "assets/battle_page.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing")
		return
	}
	err = t.ExecuteTemplate(w, "battle_page", nil)
	if err != nil {
		fmt.Println("Error executing")
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/battle":
		battle(w, r)
	default:

		fmt.Fprint(w, "Default")

	}
}
