package web

import (
	"Battleships/data"
	"Battleships/requests"
	"Battleships/web/battle"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
	"strconv"
	"time"
)

// Without this after redirect the first couple of seconds will contain no information about battle
// Enemy nickname and description would be unavailable
// Timer and current turn would be wrong
func CheckBattleDataIntegrity() {
Status:
	gameStatus, err := requests.GetGameStatus(data.GetToken())
	if err != nil || gameStatus.Opponent == "" {
		goto Status
	}
	data.SetGameStatus(gameStatus)

Data:
	enemyData, err := requests.GetEnemyData(data.GetToken())
	if err != nil {
		goto Data
	}
	data.SetEnemyData(enemyData.Nickname, enemyData.Description)

	return
}

// Max time for operation is 15 seconds, and trigger is called every ten
// So we are sending a refresh request untill no error
func RefreshLobby() {
	println("Start refresh")
Refresh:
	err := requests.GetGameRefresh(data.GetToken())
	if err != nil {
		time.Sleep(1 * time.Second)
		goto Refresh
	}
	println("Refreshed")
}

func CheckIfSomeoneJoinedLobby(c *gin.Context) {
	// When someone joins all fields in game status would be filled
	status, err := requests.GetGameStatus(data.GetToken())
	if err != nil {
		return
	}

	// Check condition using the opponent
	// If opponent is set that means the game status is valid
	if status.Opponent != "" {
		Redirect(c, "/battle")
	}
}

func JoinPlayerLobby(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	parsedData, err := url.ParseQuery(string(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Getting enemy nickname from html
	chosenLobby := parsedData.Get("chosenLobby")
	println("Chosen lobby: " + chosenLobby)

	// Starting battle passing enemy nickname and lack of bot
	err = battle.StartBattle(chosenLobby, false)
	Redirect(c, "/battle")

	if err != nil {
		println("ERROR: " + err.Error())
	}
}

// No repeated condition untill success
// If no lobbies are found user will simply click refresh again
func FindLobbies() []data.WaitingPlayer {
	servers, err := requests.GetLobby()
	if err != nil {
		return []data.WaitingPlayer{}
	}

	println("Lobbies amount: " + strconv.Itoa(len(servers)))
	return servers
}
