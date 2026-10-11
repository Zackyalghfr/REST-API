package models

import "example.com/belajar-go/config"
import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 	 string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
	Events []Event  `json:"-" gorm:"foreignKey:UserId"`
}

// Fungsi untuk menyimpan user
func (u *User) Save() error {
	return config.DB.Create(u).Error
}

// untuk mengambil semua data user
func GetAllUsers() ([]User, error) {
	var users []User

	result := config.DB.Find(&users)

	return users, result.Error
}

// mengambil user berdasarkan email, sekaligus event miliknya
func GetUserByEmail(email string) (User, error) {
	var user User
	err := config.DB.Preload("Events").Where("email = ?", email).First(&user).Error
	return user, err
}