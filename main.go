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

	router := gin.Default()

	//initialize config
	app := server.Config{Router: router}

	//routes
	app.Routes()

	requestBody := pregame.BuildPostBody()

	if err := server.PostInitGame(requestBody); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
		fmt.Println("Reconnection...")
	}

	fmt.Println("Game token is: " + client.GetToken())
	//_, coords := server.GetBoard()

	router.Run(":8080")

}
