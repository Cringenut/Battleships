package game

import (
	"Battleships/client"
	"Battleships/pregame"
	"bytes"
	"errors"
	"net/http"
)

func CheckConnection() string {
	response, err := http.Get("https://go-pjatk-server.fly.dev/swagger/index.html")
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()
	return response.Status
}

/* POST */

func InitGame() error {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	if CheckConnection() != "200 OK" {
		return errors.New("couldn't connect to game")
	}

	body := pregame.BuildPostBody()

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
	defer res.Body.Close()

	client.SetToken(res.Header.Get("x-auth-token"))
	return nil
}

/* GET */
