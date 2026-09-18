package telegram

type Config interface {
	Token() string
	ChatID() string
}
