package main

import (
	"FAQ/bootstrap"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CustomValidator for echo init api validation
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get values from env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	serverPort := os.Getenv("SERVER_PORT")

	// Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize all sections
	appContainer := bootstrap.InitializeApp(db)

	// Run echo & apis
	e := echo.New()

	// Routes APIs path
	bootstrap.SetupRoutes(e, appContainer)

	// Config CustomValidator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Run server
	fmt.Printf("🚀 Server is running on :%s\n", serverPort)
	e.Logger.Fatal(e.Start(":" + serverPort))
}
