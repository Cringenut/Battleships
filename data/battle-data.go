package data

var gameData GameData
var gameStatus *GameStatus
var IsPlayerTurn = false

func (gd *GameData) InitGameData() {
	gd.Token = ""
	gd.PlayerShips = []string{}
	gd.PlayerShots = make(map[string]bool)
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
