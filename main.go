package main

import (
	"Battleships/client"
	"Battleships/game"
	"Battleships/pregame"
	"fmt"
	"net/http"
)

func main() {

	game.InitGame()
	fmt.Println(client.GetToken())

	pregame.BuildPostBody()

	http.ListenAndServe(":8080", nil)

}
