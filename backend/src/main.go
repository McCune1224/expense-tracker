package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/expense-tracker/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetPort() string {
	port := os.Getenv("PORT")
	// Really should only be used for local development
	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf(":%s", port)
}

func GetDBString() string {
	psqlDBString := os.Getenv("DATABASE_URL")
	// Without a DB not really any point in running the API
	if psqlDBString == "" {
		panic("Failed to fetch DATABASE_URL from enviornment")
	}
	return psqlDBString
}

func main() {
	app := fiber.New()
	envPort := GetPort()
	psqlDBString := os.Getenv("DATABASE_URL")
	if psqlDBString == "" {
		panic("Failed to fetch DATABASE_URL from environment")
	}
	psqlDB, err := gorm.Open(postgres.Open(psqlDBString), &gorm.Config{
		// Might want to change this to logger.Silent for production but need users first for this to be a concern
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database, error: " + err.Error())
	}

	// appHandler := handler.New(psqlDB, &model.User{}, &model.Transaction{}, &model.Category{})
	appHandler := handler.New(psqlDB)
	appHandler.AttachRoutes(app)

	app.Listen(envPort)
}
