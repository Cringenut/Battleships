package web

import (
	"Battleships/data"
	"Battleships/requests"
	"encoding/json"
	"fmt"
)

func CheckBattleDataIntegrity() {
Battle:
	err := StartBattle()
	if err != nil {
		goto Battle
	}

Status:
	gameStatus, err := requests.GetGameStatus(data.GetToken())
	if err != nil || gameStatus.Opponent == "" {
		goto Status
	}
	data.SetGameStatus(gameStatus)

Data:
	enemyData, err := requests.GetEnemyData(data.GetToken())
	if err != nil {
		goto Data
	}
	data.SetEnemyData(enemyData.Nickname, enemyData.Description)

	return
}

func MultiplayerWaitForOpponent() error {
	body := data.GameRequestBody{
		Coords: []string{
			"A1", "A3", "B9", "C7", "D1", "D2", "D3", "D4", "D7", "E7",
			"F1", "F2", "F3", "F5", "G5", "G8", "G9", "I4", "J4", "J10",
		},
		Desc:       data.GetPlayerDescription(),
		Nick:       "New_Player",
		TargetNick: "",
		WPBot:      false,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

Token:
	token, err := requests.PostInitGame(jsonBody)
	if err != nil {
		goto Token
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
