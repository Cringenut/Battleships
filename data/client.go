package data

import (
	"fmt"
	"sync"
)

var (
	tokenMutex     sync.RWMutex // to ensure safe concurrent access to the token
	token          string
	playerShips    []string
	playerShipsSet bool
	playerShots    map[string]bool
	gameData       *GetGameStatusData
)

// SetToken sets the token.
func SetToken(t string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	token = t
}

// GetToken returns the token.
func GetToken() string {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	return token
}

// SetToken sets the token.
func SetPlayerShips(s []string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	if !playerShipsSet {
		playerShips = make([]string, len(s))
		copy(playerShips, s) // Use copy to prevent external modifications through the original slice
		playerShipsSet = true
		fmt.Println("Player ships set:", playerShips)
	}
}

// GetToken returns the token.
func GetPlayerShips() []string {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	return playerShips
}

func AppendPlayerShots(coord string, hit bool) {
	if playerShots == nil {
		playerShots = make(map[string]bool) // Initialize the map if it is nil.
	}
	playerShots[coord] = hit
	fmt.Print(playerShots)
}

func GetPlayerShots() map[string]bool {
	return playerShots
}

func SetGameData(data *GetGameStatusData) {
	gameData = data
}

func GetGameData() *GetGameStatusData {
	return gameData
}
