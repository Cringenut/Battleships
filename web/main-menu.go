package web

import (
	"Battleships/data"
	"Battleships/requests"
	"time"
)

func CheckBattleDataIntegrity() {

Data:
	time.Sleep(200 * time.Millisecond)
	enemyData, err := requests.GetEnemyData(data.GetToken())
	if err != nil {
		goto Data
	}
	data.SetEnemyData(enemyData.Nickname, enemyData.Description)

	return
}
