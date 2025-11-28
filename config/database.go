package config

import (
	"fmt"
	"log"
	"os"

	"practice-golang/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading env from system")
	}
}

func ConnectDB() {
	LoadEnv()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// AutoMigrate models
	if err := database.AutoMigrate(&models.User{}, &models.RefreshToken{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	DB = database
	log.Println("MySQL Connected!")
}
