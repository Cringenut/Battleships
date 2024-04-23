package server

import (
	"Battleships/client/data"
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
		return errors.New("couldn't connect to server")
	}

	body := []byte(`{
  "coords": [
    "A1",
    "A3",
    "B9",
    "C7",
    "D1",
    "D2",
    "D3",
    "D4",
    "D7",
    "E7",
    "F1",
    "F2",
    "F3",
    "F5",
    "G5",
    "G8",
    "G9",
    "I4",
    "J4",
    "J8"
  ],
  "desc": "My first game",
  "nick": "John_Doe",
  "target_nick": "",
  "wpbot": true
}`)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data.SetToken(res.Header.Get("x-auth-token"))
	return nil
}

/* GET */
