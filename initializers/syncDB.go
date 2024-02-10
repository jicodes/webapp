package initializers

import (
	"fmt"

	"github.com/jicodes/webapp/models"
)

func SyncDB () {
	DB.AutoMigrate(&models.User{})
	if err := DB.AutoMigrate(&models.User{}); err != nil {
			fmt.Println("Error migrating the database schema:", err)
	}
}