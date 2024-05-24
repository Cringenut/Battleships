package data

var currentGameData CurrentGameData

func (cgd *CurrentGameData) Init() {
	cgd.Token = ""
	cgd.PlayerShips = []string{}
	cgd.PlayerShots = make(map[string]bool)
}

func InitializeCurrentGameData() {
	currentGameData.Init()
}

func SetToken(token string) {
	currentGameData.Token = token
}

func GetToken() string {
	return currentGameData.Token
}

func SetPlayerShips(ships []string) {
	currentGameData.PlayerShips = ships
}
