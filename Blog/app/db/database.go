package db

import (
	"blog-backend/app/db/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	err := godotenv.Load(".env")

	if err != nil {
		return fmt.Errorf("Error loading .env file")
	}

	dsn := os.Getenv("DSN")

	if dsn == "" {
		return fmt.Errorf("Error DSN is not set")
	}

	DB, err = gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}

	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Blog{},
		&models.Like{},
		&models.Comment{},
	)

	if err != nil {
		return err
	}

	log.Println("Tables completed successfully")
	return nil
}
