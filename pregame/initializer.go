package pregame

import (
	"encoding/json"
)

type GameData struct {
	Coords      []string `json:"coords"`
	Description string   `json:"desc"`
	Nickname    string   `json:"nick"`
	TargetNick  string   `json:"target_nick"`
	Wpbot       bool     `json:"wpbot"`
}

func BuildPostBody() []byte {

	gameData := GameData{
		Coords:      []string{},
		Description: "My first server",
		Nickname:    "John_Doe",
		TargetNick:  "",
		Wpbot:       true,
	}

	//gameData.Coords = PlaceShips()
	body, _ := json.Marshal(gameData)
	return body
}
