package data

import "fmt"

var gameData GameData
var gameStatus *GameStatus
var playerShots []ShotResponse
var enemyShots []string
var IsPlayerTurn = false
var enemyData EnemyData

func (gd *GameData) InitGameData() {
	gd.Token = ""
	gd.PlayerShips = []string{}
	gd.PlayerShots = []ShotResponse{}
}

func InitializeGameData() {
	gameData.InitGameData()
}

func InitializeGameStatus() {
	gameStatus = &GameStatus{}
}

func SetToken(token string) {
	gameData.Token = token
}

func GetToken() string {
	return gameData.Token
}

func SetPlayerShips(ships []string) {
	gameData.PlayerShips = ships
}

func GetPlayerShips() []string {
	return gameData.PlayerShips
}

func SetGameStatus(status *GameStatus) {
	gameStatus = status
}

func GetGameStatus() *GameStatus {
	return gameStatus
}

func IsTurnChanged() bool {
	return gameStatus.ShouldFire != IsPlayerTurn
}

func AppendPlayerShots(coord string, res string) {
	playerShots = append(playerShots, ShotResponse{coord, res})
	fmt.Print(playerShots)
}

func SetEnemyShots(shots []string) {
	enemyShots = shots
}

func GetEnemyShots() []string {
	return enemyShots
}

func GetEnemyCellType(coord string) CellType {
	if GetToken() == "" {
		return Default
	}

	if result, found := getShotResult(playerShots, coord); found {
		if result == "hit" || result == "sunk" {
			return Hit
		} else {
			return Miss
		}
	}

	return Default
}

func GetPlayerCellType(coord string) CellType {
	if GetToken() == "" {
		return Default
	}

	if StringSliceContains(GetGameStatus().OppShots, coord) {
		if StringSliceContains(GetPlayerShips(), coord) {
			return Hit
		} else {
			return Miss
		}
	} else if StringSliceContains(GetPlayerShips(), coord) {
		return Ship
	}

	return Default
}

func getShotResult(shots []ShotResponse, coord string) (string, bool) {
	for _, shot := range shots {
		if shot.Coord == coord {
			return shot.ShotResult, true
		}
	}
	return "", false
}

func SetEnemyData(nickname, description string) {
	enemyData.Nickname = nickname
	enemyData.Description = description
}

func GetEnemyData() EnemyData {
	return enemyData
}
