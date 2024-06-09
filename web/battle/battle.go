package battle

import (
	"Battleships/data"
	"Battleships/requests"
	"Battleships/web/ships"
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
	data.SetPlayerShots([]data.ShotResponse{})
	data.SetEnemySunkShips([]string{})

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
		if data.StringSliceContains(data.GetEnemySunkShips(), coord) {
			return data.Sunk
		} else if result == "hit" {
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
	data.AppendEnemySunkShips(FindShipCells(res, coord))
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

func isCoordHit(row, col int) (bool, string) {
	coord := ships.GetCoordString(row, col)
	println("candidate: " + coord)

	for _, shot := range data.GetPlayerShots() {
		if strings.EqualFold(shot.Coord, coord) && shot.ShotResult == "hit" {
			println(true)
			return true, coord
		}
	}
	println(false)
	return false, ""
}

const size = 10

var queue []string

// Find all ship cells given one hit coordinate and the list of shots
func FindShipCells(res string, hitCoord string) []string {
	if res != "sunk" {
		return []string{}
	}

	_, _, valid := ships.GetCoordPosition(hitCoord)
	if !valid {
		fmt.Println("Invalid coordinate:", hitCoord)
		return []string{}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	foundCells := []string{hitCoord}

	queue = []string{hitCoord}
	visited := map[string]bool{hitCoord: true}

	for len(queue) > 0 {
		currentCoord := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			r, c, _ := ships.GetCoordPosition(currentCoord)
			r += dir[0]
			c += dir[1]

			if r >= 0 && r < size && c >= 0 && c < size {
				isHit, coord := isCoordHit(r, c)
				if isHit && !visited[coord] {
					foundCells = append(foundCells, coord)
					queue = append(queue, coord)
					visited[coord] = true
				}
			}
		}
	}

	return foundCells
}
