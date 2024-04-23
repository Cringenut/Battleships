package main

import (
	"Battleships/client/data"
	"Battleships/client/web"
	"Battleships/server"
	"fmt"
	"net/http"
)

func main() {

	server.InitGame()
	fmt.Println(data.GetToken())

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	// Staring server
	http.HandleFunc("/", web.Handler)
	http.ListenAndServe("localhost:8080", nil)
}
