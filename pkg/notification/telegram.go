package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Telegram struct {
	key string
}

func NewTelegramProvider(key string) *Telegram {
	return &Telegram{key: key}
}

type payload struct {
	ClientId int    `json:"chat_id"`
	Text     string `json:"text"`
}

func (t Telegram) Notify(id string, message string) error {
	clientId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	body, err := json.Marshal(&payload{
		ClientId: clientId,
		Text:     message,
	})

	if err != nil {
		return err
	}

	_, err = http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.key), "application/json", bytes.NewReader(body))

	if err != nil {
		return err
	}

	return nil
}
