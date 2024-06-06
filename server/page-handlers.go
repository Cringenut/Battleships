package server

import (
	"Battleships/data"
	"Battleships/requests"
	"Battleships/views"
	"Battleships/web"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"time"
)

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

// Abandoning the game when handle this pages for cases when user tries to go and change parameters during the battle
// Handling Main Menu page
func (app *Config) HandleMainMenu(c *gin.Context) {
	Render(c, 200, views.MakeMainMenu())
	requests.GameAbandon(data.GetToken())
}

// Handling Settings Page
func (app *Config) HandleSettings(c *gin.Context) {
	requests.GameAbandon(data.GetToken())
	web.SetCurrentSettingsPlacementType(data.GetPlayerShipPlacementType())
	Render(c, 200, views.MakeSettingsPage(data.GetPlayerNickname(), data.GetPlayerDescription()))
}

// Handling Battle Page
func (app *Config) HandleBattlePage(c *gin.Context) {
	Render(c, 200, views.MakeBattlePage())
}

// Handling Ranking Page
func (app *Config) HandleRankingPage(c *gin.Context) {
Ranking:
	ranking, err := requests.GetStats()
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		goto Ranking
	}
	Render(c, 200, views.MakeRankingPage(ranking))
}
