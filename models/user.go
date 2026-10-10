package models

import "example.com/belajar-go/config"
import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 	 string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
	Events []Event
}

// Fungsi untuk menyimpan user
func (u *User) Save() error {
	return config.DB.Create(u).Error
}