package web

import (
	"Battleships/data"
	"Battleships/requests"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
