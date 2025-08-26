package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type TelegramService struct {
	BotToken string
	ChatID   string
}

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func NewTelegramService() *TelegramService {
	return &TelegramService{
		BotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		ChatID:   os.Getenv("TELEGRAM_CHAT_ID"),
	}
}

func (t *TelegramService) SendNotification(message string) error {
	if t.BotToken == "" || t.ChatID == "" {
		return fmt.Errorf("telegram bot token or chat ID not configured")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken)

	telegramMsg := TelegramMessage{
		ChatID: t.ChatID,
		Text:   message,
	}

	jsonData, err := json.Marshal(telegramMsg)
	if err != nil {
		return fmt.Errorf("error marshaling telegram message: %v", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending telegram message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned status: %d", resp.StatusCode)
	}

	return nil
}