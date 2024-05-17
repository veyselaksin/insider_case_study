package migration

import (
	"insider_case_study/db/models"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

var onlyOnce sync.Once

func Migrate(connection *gorm.DB) {

	onlyOnce.Do(func() {

		log.Info("Migrating the database...")

		err := connection.AutoMigrate(
			models.User{},
		)
		if err != nil {
			log.Error("Error migrating the database: ", err)
		} else {
			log.Info("Database migration is successful.")
		}

		initialValues(connection)
	})

}

func initialValues(connection *gorm.DB) {

	var userData []models.User
	result := connection.Find(&userData)
	if result.Error != nil {
		log.Error("Error getting user values: ", result.Error)

	}

	if len(userData) <= 0 {
		result = connection.Create(&[]models.User{
			{
				PhoneNumber: "+905555555555",
				Content:     "Insider Project",
				Status:      "pending",
			},
			{
				PhoneNumber: "+905555555556",
				Content:     "Hello Insider",
				Status:      "pending",
			},
			{
				PhoneNumber: "+905555555557",
				Content:     "Hello World",
				Status:      "pending",
			},
			{
				PhoneNumber: "+905555555558",
				Content:     "Hello Gophers",
				Status:      "pending",
			},
		})

		if result.Error != nil {
			log.Error("Error creating user values: ", result.Error)
		}
	}
}
