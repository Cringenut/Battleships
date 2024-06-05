package web

import (
	"Battleships/data"
	"Battleships/requests"
)

func CheckBattleDataIntegrity() {

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
