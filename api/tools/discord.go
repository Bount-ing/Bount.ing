package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendDiscordMessage(channel string, message string) error {
	if message == "" {
		return nil
	}
	webhookURL := ""
	if channel == "bounties" {
		webhookURL = os.Getenv("DISCORD_BOUNTIES_WEBHOOK_URL")
	} else if channel == "errors" {
		webhookURL = os.Getenv("DISCORD_ERRORS_WEBHOOK_URL")
	} else if channel == "events" {
		webhookURL = os.Getenv("DISCORD_EVENTS_WEBHOOK_URL")
	} else {
		return nil
	}

	// Send the message to the Discord webhook

	payload := map[string]string{"content": message}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Send the HTTP request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}
