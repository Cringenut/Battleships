package data

var playerData PlayerData

func (pd *PlayerData) Init() {
	pd.Nickname = "John_Doe"
	pd.Description = "My first game"
	pd.ShipCoords = []string{}
	pd.ShipPlacementType = Simple
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

func GetPlayerShipPlacementType() PlacementType {
	return playerData.ShipPlacementType
}

func SetPlayerShipPlacementType(placement PlacementType) {
	playerData.ShipPlacementType = placement
}

func InitializePlayerData() {
	playerData.Init()
}
