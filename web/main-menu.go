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
	status, err := requests.GetGameStatus(data.GetToken())
	if err != nil {
		return
	}

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

	// Checking if next or previous type was chosen
	chosenLobby := parsedData.Get("chosenLobby")
	println("Chosen lobby: " + chosenLobby)

	err = battle.StartBattle(chosenLobby, false)
	Redirect(c, "/battle")

	if err != nil {
		println("ERROR: " + err.Error())
	}
}

func FindLobbies() []data.WaitingPlayer {
	servers, err := requests.GetLobby()
	if err != nil {
		return []data.WaitingPlayer{}
	}

	println("Lobbies amount: " + strconv.Itoa(len(servers)))
	return servers
}
