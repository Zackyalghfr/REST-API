package controllers

import (
	"net/http"
	"os"
	"time"

	"example.com/belajar-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func LoginUser(context *gin.Context) {
	var input AuthInputLogin
	// validation
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	// mengecek email apakah sudah didaftar
	user, err := models.GetUserByEmail(input.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau password salah",
		})
		return
	}


	// mengecek apakah password sesuai
	errMatchPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if errMatchPassword != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Password Salah",
		})
		return
	}

	// membuat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub" : user.ID,
	"exp" : time.Now().Add(time.Hour * 24 * 7).Unix(),
	 })

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Gagal membuat token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
			"message" : "Login Berhasil",
			"token" : tokenString,
			"user" : gin.H{
				"id"    : user.ID,
				"name"  : user.Name,
				"email" : user.Email,
				"events" : user.Events,
			},
		})

}

// mengambil data user
func GetRegister(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch events",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, users)
}