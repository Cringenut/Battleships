package data

type PlayerData struct {
	Nickname    string
	Description string
}

var playerData PlayerData

func (playerData *PlayerData) init() {
	playerData.Nickname = "John_Doe"
	playerData.Description = "My first game"
}

func GetPlayerNickname() string {
	return playerData.Nickname
}
func GetPlayerDescription() string {
	return playerData.Description
}
