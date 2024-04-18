package data

import "sync"

var (
	tokenMutex sync.RWMutex // to ensure safe concurrent access to the token
	token      string
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
