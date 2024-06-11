package web

import (
	"Battleships/data"
	"Battleships/requests"
	"time"
)

// Simply getting the leaders untill no error, because this functionality is called only once
func GetCurrentRanking() []data.PlayerStat {
Ranking:
	ranking, err := requests.GetStats()
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		goto Ranking
	}

	return ranking
}
