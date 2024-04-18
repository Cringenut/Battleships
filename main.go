package main

import (
	"Battleships/client"
	"Battleships/server"
	"fmt"
)

func main() {
	fmt.Print(server.CheckConnection())

	client.SetToken(server.PostGameDataToGetToken())
	fmt.Printf(client.GetToken())
}
