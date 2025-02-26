package processor

import (
	"fmt"
	"log"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/bount-ing/bount.ing/api/tools"
	"gorm.io/gorm"
)

func processUserCreation(message map[string]interface{}) error {
	var users []models.User

	log.Println("Processing user creation event")
	time.Sleep(3 * time.Second)

	dbc := db.DB.Find(&users)
	if dbc.Error == gorm.ErrRecordNotFound {
		log.Println("No users found")
		return fmt.Errorf("no users found: %w", dbc.Error)
	} else if dbc.Error != nil {
		log.Println("Error while fetching users:", dbc.Error)
		return fmt.Errorf("error while fetching users: %w", dbc.Error)
	}

	amountOfUsers := len(users)

	messageStr := fmt.Sprintf("New users created !\nWe are now %d users\n", amountOfUsers)

	log.Print(message)
	tools.SendDiscordMessage("events", messageStr)

	return nil
}
