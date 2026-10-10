package controllers

import (
	"net/http"

	"example.com/belajar-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthInputRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type AuthInputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(context *gin.Context) {
	// validation
	var input AuthInputRegister

	err := context.ShouldBindJSON(&input);

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	// Hash password
	hashedPassword,errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost);
	if errHash != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : "Enkripsi Password Gagal",
		})
		return
	}

	// simpan ke db
	user := models.User{
		Name: input.Name,
		Email: input.Email,
		Password: string(hashedPassword),
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Email mungkin sudah terdaftar",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
			"message" : "Register Berhasil!",
			"user" : gin.H{
				"id"    : user.ID,
				"name"  : user.Name,
				"email" : user.Email,
				"events" : user.Events,
			},
		})



}
