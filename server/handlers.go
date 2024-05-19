package server

import (
	"Battleships/data"
	"Battleships/views"
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
	render(c, 200, views.MakeMainMenu())
}

func (app *Config) HandleMainMenuContainer(c *gin.Context) {
	// Taking request body to extract chosen option
	jsonData, _ := io.ReadAll(c.Request.Body)
	// Using a variable declared inside html file using HTMX
	chosenOption := strings.TrimPrefix(string(jsonData), "chosenOption=")

	switch chosenOption {
	case "single":
		render(c, 200, views.MakeSingeplayerChosen())
	case "back":
		render(c, 200, views.MakeMainMenu())

	default:
		render(c, 200, views.MakeMainMenu())
	}

	fmt.Println(chosenOption)
}

func (app *Config) HandleBattlePageRedirect(c *gin.Context) {
	fmt.Println("Redirect")

	render(c, 200, views.MakeMainMenu())
}

// Handling Battle Page
func (app *Config) HandleBattlePage(c *gin.Context) {
	fmt.Println("Battle page")
	render(c, 200, views.MakeBattlePage())
}

// Handling Settings Page
func (app *Config) HandleSettings(c *gin.Context) {
	render(c, 200, views.MakeSettingsPage(data.GetPlayerNickname(), data.GetPlayerDescription()))
}

func (app *Config) HandleSave(c *gin.Context) {
	jsonData, _ := io.ReadAll(c.Request.Body)
	// Using a variable declared inside html file using HTMX
	saveData := strings.TrimPrefix(string(jsonData), "")

	fmt.Println("Save data:", saveData) // Add this line for debugging

	// Respond with an HTML page containing JavaScript to redirect
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Redirecting...</title>
		</head>
		<body>
			<p>Processing complete. Redirecting...</p>
			<script type="text/javascript">
				window.location.href = "/";
			</script>
		</body>
		</html>
	`)
	c.Abort() // End the request early

	fmt.Println("Redirecting to settings page")
}
