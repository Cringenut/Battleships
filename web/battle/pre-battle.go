package battle

import (
	"Battleships/data"
	"Battleships/requests"
	"Battleships/web/ships"
	"encoding/json"
	"fmt"
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
	data.SetToken("")

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
	data.SetEnemyShots([]string{})
	data.SetEnemySunkShips([]string{})
	data.SetPlayerSunkShips([]string{})

	// Setting up original positions of all ships
	// Used to show ships visually on the game board or to determine if the ship is hit
	data.SetPlayerShips(ships)
	FindPlayerShipFormations(ships)
	return nil
}

func FindAllShipCells(coord string) []string {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	foundCells := []string{coord}
	queue = []string{coord}
	visited := map[string]bool{coord: true}

	for len(queue) > 0 {
		currentCoord := queue[0]
		queue = queue[1:]
		visited[currentCoord] = true

		for _, dir := range directions {
			r, c, _ := ships.GetCoordPosition(currentCoord)
			r += dir[0]
			c += dir[1]

			if r >= 0 && r < size && c >= 0 && c < size {
				if data.StringSliceContains(data.GetPlayerShips(), ships.GetCoordString(r, c)) && !visited[ships.GetCoordString(r, c)] {
					foundCells = append(foundCells, ships.GetCoordString(r, c))
					queue = append(queue, ships.GetCoordString(r, c))
					visited[ships.GetCoordString(r, c)] = true
				}

			}
		}
	}

	return foundCells
}

func FindPlayerShipFormations(shipCoords []string) {
	visited := map[string]bool{shipCoords[0]: true}
	var foundShips [][]string

	for _, coord := range shipCoords {
		if visited[coord] {
			continue
		}

		for _, coord = range FindAllShipCells(coord) {
			visited[coord] = true
		}
		foundShips = append(foundShips, FindAllShipCells(coord))
	}

	data.SetPlayerShipsFormation(foundShips)
	fmt.Printf("%v", data.GetPlayerShipsFormation())
}
