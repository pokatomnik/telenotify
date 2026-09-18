package errors

import "errors"

// domain errors live here
var (
	// Telegram response error
	ErrorTelegramResponse = errors.New("Telegram sent error response")
)
