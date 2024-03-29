package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Message represents a Telegram message.
type Message struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

// SendMessage sends a message to given URL.
func SendMessage(url string, message *Message) error {

	fmt.Println(message.Text)
	// Cut message if its length is more than 1000 chars
	if len(message.Text) >= 1000 {
		message.Text = message.Text[:1000]
	}
	fmt.Println(message.Text)
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send successful request. Status was %q", response.Status)
	}
	return nil
}
