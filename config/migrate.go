package config

import (
	"eco-journal/entities"
	"fmt"
)

func MigrateDB() {
	if DB == nil {
		panic("Database connection not initialized")
	}

	err := DB.AutoMigrate(&entities.User{}, &entities.Journal{}, &entities.Author{}, &entities.Category{})
	if err != nil {
		fmt.Println("Failed to migrate database:", err)
		panic("Database migration failed")
	}

	fmt.Println("Database migration completed successfully")
}
