package requests

import (
	"Battleships/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func PostInitGame(body []byte) (string, error) {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	if err := CheckConnection(); err != nil {
		return "", err
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// Handle error
		}
	}(res.Body)

	return res.Header.Get("x-auth-token"), nil
}

func PostFire(token string, coord string) (string, error) {
	posturl := "https://go-pjatk-server.fly.dev/api/game/fire"

	// Creating the JSON body using a map and marshal it to bytes
	jsonData := map[string]string{"coord": coord}
	body, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	if err := CheckConnection(); err != nil {
		return "", err
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	// Adding the necessary headers
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Auth-Token", token) // Assuming authToken is provided when function is called

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Read the response body
	responseData := &data.FireResponse{}
	err = json.NewDecoder(res.Body).Decode(responseData)
	if err != nil {
		return "", err
	}

	if responseData.Result == "" {
		return "", errors.New("server response is empty")
	}

	// Print or use the "result" parameter
	fmt.Println("Response from server:", responseData.Result)
	return responseData.Result, nil
}
