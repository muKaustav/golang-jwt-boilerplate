package initializers

import (
	"golang/models"
)

func SyncDatabase() {
	if err := DB.AutoMigrate(&models.Author{}); err != nil {
		panic(err)
	}
}
