package data

type PlayerData struct {
	Nickname    string
	Description string
}

var playerData PlayerData

func (pd *PlayerData) Init() {
	pd.Nickname = "John_Doe"
	pd.Description = "My first game"
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

func InitializePlayerData() {
	playerData.Init()
}
