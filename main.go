package main

import (
	"net/http"

	"example.com/belajar-go/config"
	"example.com/belajar-go/controllers"
	"example.com/belajar-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	// menyesuaikan antara struktur struct dengan tabel di db
	config.DB.AutoMigrate(&models.Event{})

	server := gin.Default()

	// route
	api := server.Group("/api")
	{
		api.POST("/events", controllers.CreateEvents)
		api.GET("/events", controllers.GetEvents)
	}

	server.Run(":8080")
}

// function handler
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch events",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, events)
}


