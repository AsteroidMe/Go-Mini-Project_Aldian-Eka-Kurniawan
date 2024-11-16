package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		panic("failed to load env")
	}

	configDB := ConfigDB{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.User,
		configDB.Password,
		configDB.Host,
		configDB.Port,
		configDB.Name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		panic("Database not connected")
	}

	fmt.Println("Database connected successfully")
}
