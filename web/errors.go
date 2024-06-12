package web

import (
	"sync"
	"time"
)

type CustomError struct {
	Message string
	Created time.Time
	Timer   *time.Timer
}

type ErrorSlice struct {
	errors []CustomError
	mu     sync.Mutex
}

var mainMenuErrors ErrorSlice

func GetMainMenuErrors() *ErrorSlice {
	return &mainMenuErrors
}

func (es *ErrorSlice) AddError(message string) {
	es.mu.Lock()
	defer es.mu.Unlock()

	customError := CustomError{
		Message: message,
		Created: time.Now(),
		Timer:   time.NewTimer(5 * time.Second),
	}

	es.errors = append(es.errors, customError)

	// Start the timer to remove the error after 5 seconds
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

func AddMainMenuError() {
	mainMenuErrors.AddError("Test")
}
