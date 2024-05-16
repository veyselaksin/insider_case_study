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
			models.Status{},
			models.User{},
		)
		if err != nil {
			log.Error("Error migrating the database: ", err)
		} else {
			log.Info("Database migration is successful.")
		}

	})

}
