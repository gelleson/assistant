/*
 * MIT License
 *
 * Copyright (c) 2020 gelleson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Telegram is implementation of the Notificator API interface
type Telegram struct {
	key string
}

// NewTelegramProvider function to create Telegram struct with API key
func NewTelegramProvider(key string) *Telegram {
	return &Telegram{key: key}
}

type payload struct {
	ClientId int    `json:"chat_id"`
	Text     string `json:"text"`
}

// Notify is method to send message through telegram server as bot
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
