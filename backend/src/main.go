package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
    "github.com/mccune1224/expense-tracker/handler"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf(":%s", port)
}

func main() {
	app := fiber.New()
	port := GetPort()

    appHandler := handler.New()
    appHandler.AttachRoutes(app)

	app.Listen(port)
}
