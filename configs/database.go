package configs

import (
	"api-fiber-gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgreSQL() *gorm.DB {
	dsn := EnvPostgreDBUrl()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrateErr := db.AutoMigrate(&models.Book{})
	if migrateErr != nil {
		panic(migrateErr)
	}

	return db
}
