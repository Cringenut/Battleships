package server

import (
	"Battleships/views"
	"Battleships/web"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (app *Config) HandleMainMenuErrors(c *gin.Context) {
	fmt.Println("Main menu errors: %v")

	for _, _ = range web.GetMainMenuErrors().ListErrors() {
		Render(c, 200, views.MakeErrorMessage())
	}
}
