package data

var playerData PlayerData

func (pd *PlayerData) Init() {
	pd.Nickname = "Player"
	pd.Description = "My first game"
	pd.ShipCoords = []string{}
	pd.ShipPlacementType = Random
}

func SetPlayerData(nickname, description string) {
	playerData.Nickname = nickname
	playerData.Description = description
}

func GetPlayerNickname() string {
	return playerData.Nickname
}

func GetPlayerDescription() string {
	return playerData.Description
}

func GetPlayerShipPlacementType() PlacementType {
	return playerData.ShipPlacementType
}

func SetPlayerShipPlacementType(placement PlacementType) {
	playerData.ShipPlacementType = placement
}

func InitializePlayerData() {
	playerData.Init()
}
