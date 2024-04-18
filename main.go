package main

import (
	"Battleships/client/data"
	"Battleships/server"
	"fmt"
)

func main() {
	fmt.Print(server.CheckConnection())

	data.SetToken(server.PostGameDataToGetToken())
	fmt.Printf(data.GetToken())
}
