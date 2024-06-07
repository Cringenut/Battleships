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
