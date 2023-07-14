package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	jwtLogger "github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/expense-tracker/handler"
	"github.com/mccune1224/expense-tracker/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
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
		panic("FAILED TO FETCH DATABASE_URL FROM ENVIRONMENT")
	}
	return psqlDBString
}

func main() {
	app := fiber.New()
	// Sane Middlewares to use for all routes
	app.Use(jwtLogger.New())
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	envPort := GetPort()
	psqlDBString := os.Getenv("DATABASE_URL")
	if psqlDBString == "" {
		panic("FAILED TO FETCH DATABASE_URL FROM ENVIRONMENT")
	}
	psqlDB, err := gorm.Open(postgres.Open(psqlDBString), &gorm.Config{
		// Might want to change this to logger.Silent for production but need users first for this to be a concern
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic("failed to connect database, error: " + err.Error())
	}

	// migrations := []interface{}{
	// 	&model.User{},
	// 	&model.Transaction{},
	// 	&model.Category{},
	// }
	// appHandler := handler.New(psqlDB, &model.User{}, &model.Transaction{}, &model.Category{})
	psqlUserStore := store.NewPostgreUserStore(psqlDB)
	psqlTransactionStore := store.NewPostgreTransactionStore(psqlDB)
	appHandler := handler.New(psqlDB,
		psqlUserStore,
		psqlTransactionStore,
		// migrations...,
	)

	appHandler.AttachRoutes(app)

	app.Listen(envPort)
}
