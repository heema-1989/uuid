package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func LoadEnvVariables() {
	// Load .env file
	envError := godotenv.Load("./.env")
	if envError != nil {
		log.Fatal("Error loading .env file", envError)
	}
}
func ConnectToDatabase() {
	var ConnectionError error

	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, ConnectionError = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if ConnectionError != nil {
		log.Fatal("Error connecting to database ", ConnectionError)
	}
}
