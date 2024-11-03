package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func ConnectDB() *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	switch os.Getenv("SQL_DRIVER") {
	case "mysql":
		dsn = os.Getenv("MYSQL_DSN")
		if dsn == "" {
			log.Fatal("MYSQL_DSN is not set")
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	case "postgres":
		dsn = os.Getenv("POSTGRES_DSN")
		if dsn == "" {
			dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Africa/Lagos",
				os.Getenv("POSTGRES_HOST"),
				os.Getenv("POSTGRES_USER"),
				os.Getenv("POSTGRES_PASSWORD"),
				os.Getenv("POSTGRES_DB"),
				os.Getenv("POSTGRES_PORT"),
			)
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	default:
		log.Fatal("Unsupported SQL_DRIVER. Use either 'mysql' or 'postgres'")
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DBInstance = db
	log.Println("Database connection established successfully")
	return DBInstance
}
