package server

import (
	"Battleships/views"
	"Battleships/web/errors"
	"github.com/gin-gonic/gin"
)

func (app *Config) HandleMainMenuErrors(c *gin.Context) {
	for index, _ := range errors.GetMainMenuErrors().ListErrors() {
		newError := errors.GetMainMenuErrors().ListErrors()[len(errors.GetMainMenuErrors().ListErrors())-1-index]
		if newError != "" {
			Render(c, 200, views.MakeErrorMessage(newError))
		}
	}
}

func (app *Config) HandleSettingsErrors(c *gin.Context) {
	for index, _ := range errors.GetSettingsErrors().ListErrors() {
		newError := errors.GetSettingsErrors().ListErrors()[len(errors.GetSettingsErrors().ListErrors())-1-index]
		if newError != "" {
			Render(c, 200, views.MakeErrorMessage(newError))
		}
	}
}

func (app *Config) HandleBattleErrors(c *gin.Context) {
	for index, _ := range errors.GetBattleErrors().ListErrors() {
		newError := errors.GetBattleErrors().ListErrors()[len(errors.GetBattleErrors().ListErrors())-1-index]
		if newError != "" {
			Render(c, 200, views.MakeErrorMessage(newError))
		}
	}
}
