package server

import (
	"Battleships/data"
	"Battleships/views"
	"Battleships/web"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

// Handling Main Menu page
func (app *Config) HandleMainMenu(c *gin.Context) {
	Render(c, 200, views.MakeMainMenu())
}

// Handling Battle Page
func (app *Config) HandleBattlePage(c *gin.Context) {
	Render(c, 200, views.MakeBattlePage())
}

// Handling Settings Page
func (app *Config) HandleSettings(c *gin.Context) {
	web.SetCurrentSettingsPlacementType(data.GetPlayerShipPlacementType())
	Render(c, 200, views.MakeSettingsPage(data.GetPlayerNickname(), data.GetPlayerDescription()))
}
