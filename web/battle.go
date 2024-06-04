package web

import (
	"Battleships/data"
	"Battleships/requests"
	"encoding/json"
	"fmt"
)

func StartBattle() error {
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
