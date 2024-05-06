package client

import (
	"fmt"
	"sync"
)

var (
	tokenMutex  sync.RWMutex // to ensure safe concurrent access to the token
	token       string
	playerShips []string
	playerShots map[string]bool
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
	playerShips = s
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
