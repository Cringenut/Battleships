package pregame

import (
	"encoding/json"
	"fmt"
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
		Description: "My first game",
		Nickname:    "John_Doe",
		TargetNick:  "",
		Wpbot:       true,
	}

	gameData.Coords = PlaceShips()
	fmt.Println(gameData.Coords)

	body, _ := json.Marshal(gameData)
	return body
}
