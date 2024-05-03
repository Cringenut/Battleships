package pregame

import "encoding/json"

type GameData struct {
	Coords      []string `json:"coords"`
	Description string   `json:"desc"`
	Nickname    string   `json:"nick"`
	TargetNick  string   `json:"target_nick"`
	Wpbot       bool     `json:"wpbot"`
}

func BuildPostBody() []byte {

	gameData := GameData{
		Coords:      []string{"A1", "A3", "B9", "C7", "D1", "D2", "D3", "D4", "D7", "E7", "F1", "F2", "F3", "F5", "G5", "G8", "G9", "I4", "J4", "J8"},
		Description: "My first game",
		Nickname:    "John_Doe",
		TargetNick:  "",
		Wpbot:       true,
	}

	gameData.Coords = PlaceShips()

	body, _ := json.Marshal(gameData)
	return body
}
