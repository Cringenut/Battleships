package battle

import (
	"Battleships/data"
	"Battleships/requests"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
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

	data.SetShotsHistory([]data.ShotHistory{})
	data.SetPlayerAccuracy(100.0)
	data.SetEnemyAccuracy(100.0)

	// Setting up original positions of all ships
	// Used to show ships visually on the game board or to determine if the ship is hit
	data.SetPlayerShips(ships)
	return nil
}

func GetEnemyCellType(coord string) data.CellType {
	if data.GetToken() == "" {
		return data.Default
	}

	if result, found := getPlayerShotResult(data.GetPlayerShots(), coord); found {
		if result == "hit" || result == "sunk" {
			return data.Hit
		} else {
			return data.Miss
		}
	}

	return data.Default
}

func GetPlayerCellType(coord string) data.CellType {
	if data.GetToken() == "" {
		return data.Default
	}

	if data.StringSliceContains(data.GetGameStatus().OppShots, coord) {
		if data.StringSliceContains(data.GetPlayerShips(), coord) {
			return data.Hit
		} else {
			return data.Miss
		}
	} else if data.StringSliceContains(data.GetPlayerShips(), coord) {
		return data.Ship
	}

	return data.Default
}

func getPlayerShotResult(shots []data.ShotResponse, coord string) (string, bool) {
	for _, shot := range shots {
		if shot.Coord == coord {
			return shot.ShotResult, true
		}
	}
	return "", false
}

func FireAtEnemy(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	coord := strings.TrimPrefix(string(jsonData), "coord=")
	res, err := requests.PostFire(data.GetToken(), coord)
	if err != nil {
		return
	}

	data.AppendPlayerShots(coord, res)
	data.AppendShotsHistory(coord, res, data.GetPlayerNickname())
}

func CalculateEnemyAccuracy() {
	if len(data.GetEnemyShots()) == 0 {
		data.SetEnemyAccuracy(100.00)
		return
	}

	hit := 0
	for _, shot := range data.GetEnemyShots() {
		if data.StringSliceContains(data.GetPlayerShips(), shot) {
			hit++
		}
	}

	accuracy := (float64(hit) / float64(len(data.GetEnemyShots()))) * 100
	data.SetEnemyAccuracy(accuracy)
}

func CalculatePlayerAccuracy() {
	if len(data.GetPlayerShots()) == 0 {
		data.SetPlayerAccuracy(100.00)
		return
	}

	hit := 0
	for _, shot := range data.GetPlayerShots() {
		if shot.ShotResult != "miss" {
			hit++
		}
	}

	accuracy := (float64(hit) / float64(len(data.GetPlayerShots()))) * 100
	data.SetPlayerAccuracy(accuracy)
}
