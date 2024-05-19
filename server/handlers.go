package server

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

// Handling Main Menu page
func (app *Config) HandleMainMenu(c *gin.Context) {
	if c.Request.URL.Path != "/" {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	}
	//render(c, 200, views.MakeMainMenu())
}

func (app *Config) HandleMainMenuContainer(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	// Using a variable declared inside html file using HTMX
	chosenOption := strings.TrimPrefix(string(jsonData), "chosenOption=")

	switch chosenOption {
	case "single":
		//render(c, 200, views.MakeSingeplayerChosen())
	case "back":
		//render(c, 200, views.MakeMainMenu())

	default:
		//render(c, 200, views.MakeMainMenu())
	}

	fmt.Println(chosenOption)
}

func (app *Config) HandleBattlePageRedirect(c *gin.Context) {
	fmt.Println("Redirect")
	//render(c, 200, views.MakeMainMenu())
}

// Handle Battle Page
func (app *Config) HandleBattlePage(c *gin.Context) {
	fmt.Println("Battle page")
	//render(c, 200, views.MakeBattlePage())
}
