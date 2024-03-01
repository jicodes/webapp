package initializers

import (
	"fmt"

	"github.com/jicodes/webapp/models"
)

func SyncDB () {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
			fmt.Println("Error migrating the database schema:", err)
	}
}