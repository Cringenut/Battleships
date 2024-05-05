package main

import (
	"Battleships/client"
	"Battleships/pregame"
	"Battleships/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	requestBody := pregame.BuildPostBody()

	if err := server.InitGame(requestBody); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
		fmt.Println("Reconnection...")
	}

	fmt.Println("Game token is: " + client.GetToken())
	//_, coords := server.GetBoard()

	r := gin.Default()

	// Handlers
	r.Handle("GET", "/", server.HandleHomePage())
	r.Handle("POST", "/fire", server.HandleFire())

	r.Run(":8080")

}
