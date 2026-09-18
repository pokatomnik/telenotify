package telegram

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"

	"github.com/pokatomnik/telenotify/internal/entities/errors"
	"github.com/pokatomnik/telenotify/internal/entities/notifreq"
)

type Telegram struct {
	config     Config
	httpClient http.Client
}

func New(config Config, client http.Client) Telegram {
	return Telegram{
		config:     config,
		httpClient: client,
	}
}

func (t Telegram) SendMessage(context context.Context, msg notifreq.NotificationRequest) error {
	body := ReqBody{
		ChatID: t.config.ChatID(),
		Text:   msg.NotificationText,
	}
	jsonBytes, err := json.Marshal(&body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		context,
		"POST",
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.config.Token()),
		bytes.NewBuffer(jsonBytes),
	)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := t.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: status code: %d", errors.ErrorTelegramResponse, res.StatusCode)
	}

	if _, err = io.ReadAll(res.Body); err != nil {
		return err
	}

	return nil
}
