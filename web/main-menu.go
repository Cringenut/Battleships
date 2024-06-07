package web

import (
	"Battleships/data"
	"Battleships/requests"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
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

func RedirectToBattle(c *gin.Context) {
	// Battle page we redirect tp
	targetURL := "/battle"

	time.Sleep(1000 * time.Millisecond)

	// Without this after redirect the first couple of seconds will contain no information about battle
	// Enemy nickname and description would be unavailable
	// Timer and current turn would be wrong
	CheckBattleDataIntegrity()

	// Log the redirection attempt
	log.Printf("Redirecting to: %s", targetURL)

	// Check if the request is coming from HTMX
	if c.GetHeader("HX-Request") != "" {
		// Respond with an HTMX-specific header to trigger a client-side redirect
		c.Header("HX-Redirect", targetURL)
		c.Status(http.StatusOK)
	} else {
		// Regular redirection for non-HTMX requests
		c.Redirect(http.StatusFound, targetURL)
	}
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
		RedirectToBattle(c)
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

	err = StartBattle(chosenLobby, false)
	RedirectToBattle(c)

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
