package main

import (
	"Battleships/client"
	"Battleships/game"
	"Battleships/pregame"
	"fmt"
	"net/http"
	"time"
)

func main() {

	requestBody := pregame.BuildPostBody()

	if err := game.InitGame(requestBody); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
		fmt.Println("Reconnection...")
	}

	fmt.Println("Game token is: " + client.GetToken())
	err, coords := game.GetBoard()
	client.SetShips(coords)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
