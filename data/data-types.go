package data

const (
	Random Placement = iota
	Simple
	Advanced
)

type PlayerData struct {
	Nickname          string
	Description       string
	ShipCoords        []string
	ShipPlacementType Placement
}

type CurrentGameData struct {
	Token       string
	PlayerShips []string
	PlayerShots map[string]bool
}

type ServerGameStatusData struct {
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}

type GameRequestBody struct {
	Coords     []string `json:"coords"`
	Desc       string   `json:"desc"`
	Nick       string   `json:"nick"`
	TargetNick string   `json:"target_nick"`
	WPBot      bool     `json:"wpbot"`
}

type BoardResponse struct {
	Board []string `json:"board"`
}
