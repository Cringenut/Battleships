package web

import (
	"fmt"
	"net/http"
)

type PageData struct {
	Token string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/battle":
		battlePageHandler(w, r)
	default:

		fmt.Fprint(w, "Default")
	}
}

func battlePageHandler(w http.ResponseWriter, r *http.Request) {

}
