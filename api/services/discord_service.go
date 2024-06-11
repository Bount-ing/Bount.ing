package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"open-bounties-api/models"
	"os"
)

type DiscordService struct {
}

// NewBountyService creates a new BountyService
func NewDiscordService() *DiscordService {
	return &DiscordService{}
}

func (ds *DiscordService) SendBountyCreationNotification(bounty models.Bounty, issue models.Issue, user models.User) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL_BOUNTIES")
	log.Printf("Sending notification to Discord channel for new bounty %d", bounty.ID)

	// Create the message payload
	messagePayload := map[string]interface{}{
		"username": "Oracle@Bount.ing",
		"embeds": []map[string]interface{}{
			{
				"title":       fmt.Sprintf("A new mission is available"),
				"description": issue.Title,
				"color":       0x00FF00, // Green color
				"fields": []map[string]string{
					{
						"name":  "Created By",
						"value": user.Username,
					},
					{
						"name":  "Reward",
						"value": fmt.Sprintf("%d €", bounty.Amount),
					},
					{
						"name":  "Issue URL",
						"value": issue.GithubURL,
					},
				},
			},
		},
	}

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(messagePayload)
	if err != nil {
		log.Fatalf("Error marshaling payload: %v", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Printf("Received non-OK response from Discord: %s", resp.Status)
	} else {
		log.Println("Notification sent successfully!")
	}
}
