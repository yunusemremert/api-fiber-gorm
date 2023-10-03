package migrations

import (
	"api-fiber-gorm/models"
	"gorm.io/gorm"
)

type PostgreDB struct {
	DB *gorm.DB
}

func (d PostgreDB) MigrateBookTable() {
	err := d.DB.AutoMigrate(&models.Book{})
	if err != nil {
		return
	}
}
