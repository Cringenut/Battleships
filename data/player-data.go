package data

type Placement int

var playerData PlayerData

func (pd *PlayerData) Init() {
	pd.Nickname = "John_Doe"
	pd.Description = "My first game"
	pd.ShipCoords = []string{}
	pd.ShipPlacementType = Random
}

func SetPlayerData(nickname, description string, shipCoords []string) {
	playerData.Nickname = nickname
	playerData.Description = description
	playerData.ShipCoords = shipCoords
}

func GetPlayerNickname() string {
	return playerData.Nickname
}

func GetPlayerDescription() string {
	return playerData.Description
}

func GetPlayerShipPlacementType() Placement {
	return playerData.ShipPlacementType
}

func InitializePlayerData() {
	playerData.Init()
}
