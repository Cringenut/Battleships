package battle

import (
	"Battleships/data"
	"Battleships/requests"
	"time"
)

func CheckGameStatus() {
	if data.GetToken() == "" {
		return
	}

	time.Sleep(100 * time.Millisecond)
	gameStatus, err := requests.GetGameStatus(data.GetToken())
	if err == nil {
		data.SetGameStatus(gameStatus)
	}
}

func CheckWin() bool {
	println("Battle Ended")
	if data.GetGameStatus().LastGameStatus == "win" {
		return true
	} else if data.GetGameStatus().LastGameStatus == "lose" {
		return false
	}
	return false
}

func CheckEnemyShots() {
	for _, shot := range data.GetGameStatus().OppShots[len(data.GetEnemyShots()):] {

		if !data.StringSliceContains(data.GetPlayerShips(), shot) {
			data.AppendShotsHistory(shot, "miss", data.GetEnemyData().Nickname)
			goto Set
		} else {
			var hitShip []string
			for _, playerShip := range data.GetPlayerShipsFormation() {
				if data.StringSliceContains(playerShip, shot) {
					hitShip = playerShip
					break
				}
			}

			for _, playerCoords := range hitShip {
				if !data.StringSliceContains(data.GetGameStatus().OppShots, playerCoords) {
					data.AppendShotsHistory(shot, "hit", data.GetEnemyData().Nickname)
					goto Set
				}
			}

			data.AppendShotsHistory(shot, "sunk", data.GetEnemyData().Nickname)
			data.AppendPlayerSunkShips(hitShip)
		}

	Set:
		data.SetEnemyShots(data.GetGameStatus().OppShots)
	}
}
