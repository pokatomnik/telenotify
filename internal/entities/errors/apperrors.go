package errors

import "errors"

// domain errors live here
var (
	// Telegram response error
	ErrorTelegramResponse = errors.New("Telegram sent error response")

	// Skills errors
	ErrorSkillPathIsNotAFile = errors.New("Skill path is not a file")
)
