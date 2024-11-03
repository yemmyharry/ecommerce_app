package main

import (
	"ecommerce/internal/core/domain"
	"ecommerce/internal/database"
	"ecommerce/internal/logger"

	"github.com/joho/godotenv"
)

func Migrate() {

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	db := database.ConnectDB()
	db.AutoMigrate(&domain.User{}, &domain.Order{}, &domain.Product{})
}

func main() {
	Migrate()
}
