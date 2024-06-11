package battle

import (
	"Battleships/data"
	"Battleships/requests"
	"Battleships/web/ships"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

func GetEnemyCellType(coord string) data.CellType {
	if data.GetToken() == "" {
		return data.Default
	}

	// If cell was shot find the shot result
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

	if data.StringSliceContains(data.GetEnemyShots(), coord) {
		if data.StringSliceContains(data.GetPlayerSunkShips(), coord) {
			return data.Sunk
		} else if data.StringSliceContains(data.GetPlayerShips(), coord) {
			return data.Hit
		} else {
			return data.Miss
		}
		// We know the player ship placement so we can show it
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

	// Append player shots history
	data.AppendPlayerShots(coord, res)
	data.AppendShotsHistory(coord, res, data.GetPlayerNickname())
	// Checks if player shot was a sunk and if so finds sunken ship cells
	data.AppendEnemySunkShips(FindEnemyShipCells(res, coord))
}

// Getting amount of shots
// Counting hits
// Calculating accuracy
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

const size = 10

var queue []string

// Checking if "candidate" is an enemy hit cell
func isEnemyCoordHit(row, col int) (bool, string) {
	coord := ships.GetCoordString(row, col)
	for _, shot := range data.GetPlayerShots() {
		// Appending if shot is among player shots and resulted into hit
		if strings.EqualFold(shot.Coord, coord) && shot.ShotResult == "hit" {
			println(true)
			return true, coord
		}
	}
	println(false)
	return false, ""
}

// Find all ship cells given one hit coordinate and the list of shots
func FindEnemyShipCells(res string, hitCoord string) []string {
	// Not sunk return
	if res != "sunk" {
		return []string{}
	}

	// Checking if valid coord was passed
	_, _, valid := ships.GetCoordPosition(hitCoord)
	if !valid {
		return []string{}
	}

	// left, right, down, up
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// Adding the sunk coord as first
	foundCells := []string{hitCoord}

	// Coords we will iterate through
	queue = []string{hitCoord}
	// In order to don't check already checked texture and not go into infinite loop of searches
	visited := map[string]bool{hitCoord: true}

	for len(queue) > 0 {
		currentCoord := queue[0]
		queue = queue[1:]

		// Checking each direction for potential hit cell
		for _, dir := range directions {
			r, c, _ := ships.GetCoordPosition(currentCoord)
			r += dir[0]
			c += dir[1]

			if r >= 0 && r < size && c >= 0 && c < size {
				isHit, coord := isEnemyCoordHit(r, c)
				if isHit && !visited[coord] {
					foundCells = append(foundCells, coord)
					queue = append(queue, coord)
					visited[coord] = true
				}
			}
		}
	}

	// Returning all enemy ship cells
	return foundCells
}
