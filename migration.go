package main

import (
	"api-fiber-gorm/configs"
	"api-fiber-gorm/migrations"
)

func main() {
	dbClient := configs.ConnectPostgreSQL()

	rm := migrations.PostgreDB{DB: dbClient}

	rm.MigrateBookTable()
	// Other migrations
}
