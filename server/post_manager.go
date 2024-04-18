package server

import (
	"bytes"
	"net/http"
)

func PostGameDataToGetToken() string {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

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
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return res.Header.Get("x-auth-token")
}
