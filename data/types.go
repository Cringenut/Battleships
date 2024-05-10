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

type GetPlayersData struct {
	PlayerDesc string `json:"desc"`
	PlayerNick string `json:"nick"`
	EnemyDesc  string `json:"opp_desc"`
	EnemyNick  string `json:"opponent"`
}

type GameStartData struct {
	Coords      []string `json:"coords"`
	Description string   `json:"desc"`
	Nickname    string   `json:"nick"`
	TargetNick  string   `json:"target_nick"`
	Wpbot       bool     `json:"wpbot"`
}
