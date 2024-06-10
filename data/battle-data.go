package data

var gameData GameData
var gameStatus *GameStatus
var playerShots []ShotResponse
var enemyShots []string
var enemyData EnemyData
var shotsHistory []ShotHistory
var playerAccuracy = 100.0
var enemyAccuracy = 100.0
var enemySunkShips []string
var playerSunkShips []string
var playerShipsFormation [][]string

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

func AppendPlayerShots(coord string, res string) {
	playerShots = append(playerShots, ShotResponse{coord, res})
}

func GetPlayerShots() []ShotResponse {
	return playerShots
}

func SetPlayerShots(shots []ShotResponse) {
	playerShots = shots
}

func SetEnemyShots(shots []string) {
	enemyShots = shots
}

func GetEnemyShots() []string {
	return enemyShots
}

func SetEnemyData(nickname, description string) {
	enemyData.Nickname = nickname
	enemyData.Description = description
}

func GetEnemyData() EnemyData {
	return enemyData
}

func SetEnemyAccuracy(newAccuracy float64) {
	enemyAccuracy = newAccuracy
}

func GetEnemyAccuracy() float64 {
	return enemyAccuracy
}

func SetPlayerAccuracy(newAccuracy float64) {
	playerAccuracy = newAccuracy
}

func GetPlayerAccuracy() float64 {
	return playerAccuracy
}

func SetShotsHistory(newShotsHistory []ShotHistory) {
	shotsHistory = newShotsHistory
}

func AppendShotsHistory(coord string, res string, owner string) {
	shotsHistory = append(shotsHistory, ShotHistory{ShotResponse{coord, res}, owner})
}

func GetShotsHistory() []ShotHistory {
	return shotsHistory
}

func SetEnemySunkShips(ships []string) {
	enemySunkShips = ships
}

func AppendEnemySunkShips(ships []string) {
	enemySunkShips = append(enemySunkShips, ships...)
}

func GetEnemySunkShips() []string {
	return enemySunkShips
}

func SetPlayerSunkShips(ships []string) {
	playerSunkShips = ships
}

func AppendPlayerSunkShips(ships []string) {
	playerSunkShips = append(playerSunkShips, ships...)
}

func GetPlayerSunkShips() []string {
	return playerSunkShips
}

func SetPlayerShipsFormation(formation [][]string) {
	playerShipsFormation = formation
}

func GetPlayerShipsFormation() [][]string {
	return playerShipsFormation
}
