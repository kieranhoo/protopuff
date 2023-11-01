package database

import (
	"protopuff/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn, _ := utils.ConnectionURLBuilder("postgress")
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
