package telegram

type ReqBody struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}
