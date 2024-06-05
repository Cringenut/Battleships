package web

import (
	"Battleships/data"
	"Battleships/requests"
	"encoding/json"
	"fmt"
)

// Setting up data for the battle
func StartBattle() error {
	// Body that will be sent to the server to start the battle
	body := data.GameRequestBody{
		Coords: []string{
			"A1", "A3", "B9", "C7", "D1", "D2", "D3", "D4", "D7", "E7",
			"F1", "F2", "F3", "F5", "G5", "G8", "G9", "I4", "J4", "J8",
		},
		Desc:       data.GetPlayerDescription(),
		Nick:       data.GetPlayerNickname(),
		TargetNick: "",
		WPBot:      true,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

Token:
	// If request is failed try to start the game until successful
	token, err := requests.PostInitGame(jsonBody)
	if err != nil {
		goto Token
	}

	data.SetToken(token)
	// Printing token for debug
	println(data.GetToken())

Ships:
	// Gets the ships from the server
	// Not reliable because when ship is hit, it gets removed from the request body
	ships, _ := requests.GetBoard(data.GetToken())
	if len(ships) == 0 {
		goto Ships
	}
	// Printing ships for debug
	for _, position := range ships {
		fmt.Println(position)
	}

	// Setting up original positions of all ships
	// Used to show ships visually on the game board or to determine if the ship is hit
	data.SetPlayerShips(ships)
	return nil
}

// Temporary solution
// Test body and simmilar as Start battle, but uses enemy nickname as parameter
func JoinLobby(enemyNickname string) error {
	body := data.GameRequestBody{
		Coords: []string{
			"A1", "A3", "B9", "C7", "D1", "D2", "D3", "D4", "D7", "E7",
			"F1", "F2", "F3", "F5", "G5", "G8", "G9", "I4", "J4", "J8",
		},
		Desc:       data.GetPlayerDescription(),
		Nick:       data.GetPlayerNickname(),
		TargetNick: enemyNickname,
		WPBot:      false,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	token, err := requests.PostInitGame(jsonBody)
	if err != nil {
		return err
	}
	data.SetToken(token)
	println(data.GetToken())

Ships:
	ships, _ := requests.GetBoard(data.GetToken())
	if len(ships) == 0 {
		goto Ships
	}
	for _, position := range ships {
		fmt.Println(position)
	}

	data.SetPlayerShips(ships)
	return nil

}
