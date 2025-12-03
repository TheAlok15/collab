package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/TheAlok15/collab/internal/models"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database",err)
		os.Exit(1)
	}

	DB = db

	err = DB.AutoMigrate(&models.User{})
	if err != nil{
		log.Fatal("Database migration failed:", err)
	}

	fmt.Println("Database connected and user model migrated")
}