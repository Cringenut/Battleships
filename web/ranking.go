package web

import (
	"Battleships/data"
	"Battleships/requests"
	"time"
)

func GetCurrentRanking() []data.PlayerStat {
Ranking:
	ranking, err := requests.GetStats()
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		goto Ranking
	}

	return ranking
}
