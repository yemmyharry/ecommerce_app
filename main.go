package main

import (
	adapter "ecommerce/internal/adapter/api/controller"
	mysql_repo "ecommerce/internal/adapter/repositories/mysql"

	"ecommerce/internal/database"
	"ecommerce/internal/logger"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	mysql_repo.DB = database.ConnectDB()

	router := gin.Default()
	handler := adapter.NewHTTPHandler()
	handler.Routes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	logger.Info(fmt.Sprintf(" Starting server on port %v", port))
	router.Run(":" + port)
}
