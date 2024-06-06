package requests

import (
	"Battleships/data"
	"encoding/json"
	"errors"
	"fmt"
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

func GetBoard(token string) ([]string, error) {
	geturl := "https://go-pjatk-server.fly.dev/api/game/board"

	if err := CheckConnection(); err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}

	// Add the X-Auth-Token header using the token received during game initialization
	req.Header.Add("X-Auth-Token", token)

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the game board")
	}

	var boardResponse data.BoardResponse
	if err := json.NewDecoder(res.Body).Decode(&boardResponse); err != nil {
		return nil, err
	}

	return boardResponse.Board, nil
}

func GetGameStatus(token string) (*data.GameStatus, error) {
	geturl := "https://go-pjatk-server.fly.dev/api/game"

	if err := CheckConnection(); err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}

	// Add the X-Auth-Token header using the token received during game initialization
	req.Header.Add("X-Auth-Token", token)

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the game status")
	}

	var gameStatus data.GameStatus
	if err := json.NewDecoder(res.Body).Decode(&gameStatus); err != nil {
		return nil, err
	}

	// Check if game_status is empty
	if gameStatus.GameStatus == "" {
		return nil, errors.New("game status is empty")
	}

	return &gameStatus, nil
}

func GetEnemyData(token string) (*data.EnemyData, error) {
	geturl := "https://go-pjatk-server.fly.dev/api/game/desc"

	if err := CheckConnection(); err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}

	// Add the X-Auth-Token header using the token received during game initialization
	req.Header.Add("X-Auth-Token", token)

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the game status")
	}

	var enemyData data.EnemyData
	if err := json.NewDecoder(res.Body).Decode(&enemyData); err != nil {
		return nil, err
	}

	// Check if game_status is empty
	if enemyData.Nickname == "" {
		return nil, errors.New("enemy data is empty")
	}

	return &enemyData, nil
}

func GetGameRefresh(token string) error {
	geturl := "https://go-pjatk-server.fly.dev/api/game/refresh"

	if err := CheckConnection(); err != nil {
		return err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return err
	}

	// Add the X-Auth-Token header using the token received during game initialization
	req.Header.Add("X-Auth-Token", token)

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to refresh the game")
	}

	return nil
}

func GetLobby() ([]data.WaitingPlayer, error) {
	geturl := "https://go-pjatk-server.fly.dev/api/lobby"

	if err := CheckConnection(); err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get the lobbies")
	}

	var playerLobbies []data.WaitingPlayer
	if err := json.NewDecoder(res.Body).Decode(&playerLobbies); err != nil {
		return nil, err
	}

	return playerLobbies, nil
}

func GetStats() ([]data.PlayerStat, error) {
	geturl := "https://go-pjatk-server.fly.dev/api/stats"

	if err := CheckConnection(); err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client and send the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get the lobbies")
	}

	var statsResponse data.StatsResponse
	if err := json.NewDecoder(res.Body).Decode(&statsResponse); err != nil {
		return nil, err
	}

	return statsResponse.Stats, nil
}
