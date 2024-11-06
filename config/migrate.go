package config

import (
	"rest/repo/auth"

	"gorm.io/gorm"
)

func MigrateDB(DB *gorm.DB) {
	db.AutoMigrate(&auth.User{})
}
