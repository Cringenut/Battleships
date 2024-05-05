package server

import (
	"Battleships/client"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CheckConnection() error {
	response, err := http.Get("https://go-pjatk-server.fly.dev/swagger/index.html")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Connection not available")
	}

	return nil
}

/* POST */
func InitGame(body []byte) error {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	if err := CheckConnection(); err != nil {
		return err
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer((body)))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	// If post goes successful
	client.SetToken(res.Header.Get("x-auth-token"))
	return nil
}

type GetBoardResponse struct {
	ShipCoords []string `json:"board"`
}

func TestFire(c *gin.Context) {
	fmt.Println("Fired")
}

/* GET */
func GetBoard() (error, []string) {
	// Define the URL and the custom headers
	url := "https://go-pjatk-server.fly.dev/api/game/board"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err, nil
	}

	// Set the headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", client.GetToken())

	// Create a c to send the request
	c := &http.Client{}

	// Send the request
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err, nil
	}
	defer resp.Body.Close()

	var get GetBoardResponse

	// Read and display the response
	if err := json.NewDecoder(resp.Body).Decode(&get); err != nil {
		print(err)
	}

	return nil, get.ShipCoords
}
