package data

type FireResponseData struct {
	Result string `json:"result"`
}

type GetBoardResponseData struct {
	Board []string `json:"board"`
}

type GetGameStatusData struct {
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}
