package models

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	UserId      int    `json:"userid"`
	Datetime	time.Time	`json:"datetime" binding:"required"`
}

var db *gorm.DB

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db = d
}

func init() {
	initDB()

	err := db.AutoMigrate(&Event{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

// Fungsi untuk menyimpan event
func (e *Event) Save() error {
	result := db.Create(&e)
	return result.Error
}

// Fungsi menampilkan semua event
func GetAllEvents() ([]Event, error) {
	var events []Event

	result := db.Find(&events)

	return events, result.Error
}