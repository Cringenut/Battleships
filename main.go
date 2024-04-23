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
	fmt.Printf(data.GetToken())
	http.HandleFunc("/", web.WelcomePage)
	http.ListenAndServe("", nil)
}
