package errors

import (
	"sync"
	"time"
)

const lifeTime = 3 * time.Second

type CustomError struct {
	Message string
	Created time.Time
	Timer   *time.Timer
}

type ErrorSlice struct {
	errors []CustomError
	mu     sync.Mutex
}

// Slices of errors
var mainMenuErrors ErrorSlice
var settingsErrors ErrorSlice
var battleErrors ErrorSlice

func GetMainMenuErrors() *ErrorSlice {
	return &mainMenuErrors
}

func GetSettingsErrors() *ErrorSlice {
	return &settingsErrors
}

func GetBattleErrors() *ErrorSlice {
	return &battleErrors
}

func (es *ErrorSlice) AddError(message string) {
	es.mu.Lock()
	defer es.mu.Unlock()

	customError := CustomError{
		Message: message,
		Timer:   time.NewTimer(lifeTime),
	}

	es.errors = append(es.errors, customError)

	// Start the timer to remove the error
	go func(err CustomError) {
		<-err.Timer.C
		es.removeError(err)
	}(customError)
}

func (es *ErrorSlice) removeError(err CustomError) {
	es.mu.Lock()
	defer es.mu.Unlock()

	for i, e := range es.errors {
		if e == err {
			es.errors = append(es.errors[:i], es.errors[i+1:]...)
			break
		}
	}
}

func (es *ErrorSlice) ListErrors() []string {
	es.mu.Lock()
	defer es.mu.Unlock()

	errors := make([]string, len(es.errors))
	for _, err := range es.errors {
		errors = append(errors, err.Message)
	}

	return errors
}

func AddMainMenuError(errorMessage string) {
	mainMenuErrors.AddError(errorMessage)
}

func AddSettingsError(errorMessage string) {
	settingsErrors.AddError(errorMessage)
}

func AddBattleError(errorMessage string) {
	settingsErrors.AddError(errorMessage)
}
