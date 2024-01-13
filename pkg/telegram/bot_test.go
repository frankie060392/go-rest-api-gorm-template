package telegram_test

import (
	"frankie060392/rest-api-clean-arch/bootstrap"
	"frankie060392/rest-api-clean-arch/pkg/telegram"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot(t *testing.T) {
	assert := assert.New(t)
	config := bootstrap.LoadConfig(".")
	err := telegram.SendMessage(config.TelegramUrl, &telegram.Message{ChatID: config.TelegramChatId, Text: "Test112313213123123"})
	assert.Nil(err, "Equal send telegram")
}
