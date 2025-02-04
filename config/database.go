package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	env := os.Getenv("ENV")
	err := godotenv.Load("./env/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_DATABASE")
	dbPort := 25432

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, errDb := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDb != nil {
		log.Fatal("Error connect database")
	}
	return db
}
