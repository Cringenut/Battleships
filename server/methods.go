package server

import (
	"Battleships/client"
	"bytes"
	"errors"
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

func Fire(body []byte) error {
	//posturl := "https://go-pjatk-server.fly.dev/api/game/fire"

	return nil
}
