package web

import (
	"Battleships/data"
	"Battleships/requests"
	"encoding/json"
	"fmt"
	"time"
)

func bodyBuilder(enemyNickname string, isSingleplayer bool) data.GameRequestBody {
	// Body that will be sent to the server to start the battle
	body := data.GameRequestBody{
		Coords:     data.GetPlayerShips(),
		Desc:       data.GetPlayerDescription(),
		Nick:       data.GetPlayerNickname(),
		TargetNick: enemyNickname,
		WPBot:      isSingleplayer,
	}

	return body
}

// Setting up data for the battle
func StartBattle(enemyNickname string, isSingleplayer bool) error {
	jsonBody, err := json.Marshal(bodyBuilder(enemyNickname, isSingleplayer))
	if err != nil {
		return err
	}

Token:
	// If request is failed try to start the game until successful
	token, err := requests.PostInitGame(jsonBody)
	if err != nil {
		time.Sleep(200 * time.Millisecond)
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
