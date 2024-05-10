package pregame

import (
	"Battleships/data"
	"encoding/json"
)

func BuildPostBody() []byte {

	gameData := data.GameStartData{
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
