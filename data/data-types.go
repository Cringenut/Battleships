package data

type PlacementType int

const (
	Random PlacementType = iota
	Simple
	Advanced
	ServerRandom
)

type CellType int

const (
	Default CellType = iota
	Ship
	Hit
	Miss
)

type PlayerData struct {
	Nickname          string
	Description       string
	ShipCoords        []string
	ShipPlacementType PlacementType
}

type EnemyData struct {
	Nickname    string `json:"opponent"`
	Description string `json:"opp_desc"`
}

type GameData struct {
	Token       string
	PlayerShips []string
	PlayerShots []ShotResponse
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

type GameStatus struct {
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}

type FireResponse struct {
	Result string `json:"result"`
}

type ShotResponse struct {
	Coord      string
	ShotResult string
}

type WaitingPlayer struct {
	Nick       string `json:"nick"`
	GameStatus string `json:"game_status"`
}

type PlayerStat struct {
	Games  int    `json:"games"`
	Nick   string `json:"nick"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
	Wins   int    `json:"wins"`
}

type StatsResponse struct {
	Stats []PlayerStat `json:"stats"`
}
