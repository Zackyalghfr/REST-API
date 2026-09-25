package models

import (
	"example.com/belajar-go/config"
	"time"
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

// Fungsi untuk menyimpan event
func (e *Event) Save() error {
	result := config.DB.Create(&e)
	return result.Error
}

// Fungsi menampilkan semua event
func GetAllEvents() ([]Event, error) {
	var events []Event

	result := config.DB.Find(&events)

	return events, result.Error
}