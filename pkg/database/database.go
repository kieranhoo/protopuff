package database

import (
	"protopuff/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn, _ := utils.ConnectionURLBuilder("mysql")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
