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

// Used to append enemy shots to history and set them
func CheckEnemyShots() {
	// Going through difference between shots from server and that are already set
	for index, shot := range data.GetGameStatus().OppShots[len(data.GetEnemyShots()):] {

		// If not player ship
		if !data.StringSliceContains(data.GetPlayerShips(), shot) {
			data.AppendShotsHistory(shot, "miss", data.GetEnemyData().Nickname)
			goto Set
		} else {
			var hitShip []string
			// Find the ship that was hit by finding one with provided coordinate
			for _, playerShip := range data.GetPlayerShipsFormation() {
				if data.StringSliceContains(playerShip, shot) {
					hitShip = playerShip
					break
				}
			}

			// Checking all ship coords
			for _, playerCoords := range hitShip {
				// If any of cells wasn't hit yet append as hit
				// Check the length as GetEnemyShots length + current index
				// Otherwise sunk can be added multiple times
				if !data.StringSliceContains(data.GetGameStatus().OppShots[:len(data.GetEnemyShots())+index-1], playerCoords) {
					data.AppendShotsHistory(shot, "hit", data.GetEnemyData().Nickname)
					goto Set
				}
			}

			// Append sunk if all is checked true
			data.AppendShotsHistory(shot, "sunk", data.GetEnemyData().Nickname)
			data.AppendPlayerSunkShips(hitShip)
		}

		// Tag to skip all other logic if necessary
	Set:
		data.SetEnemyShots(data.GetGameStatus().OppShots)
	}
}
