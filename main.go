package main

import (
	"Battleships/assets"
	"Battleships/client/data"
	"Battleships/server"
	"fmt"
	"github.com/a-h/templ"
	"net/http"
)

func main() {

	server.InitGame()
	fmt.Println(data.GetToken())

	coords := []string{
		"A1",
		"A3",
		"B9",
		"C7",
		"D1",
		"D2",
		"D3",
		"D4",
		"D7",
		"E7",
		"F1",
		"F2",
		"F3",
		"F5",
		"G5",
		"G8",
		"G9",
		"I4",
		"J4",
		"J8",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(assets.Page(data.GetToken(), coords)).ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)

}
