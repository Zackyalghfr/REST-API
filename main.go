package main

import (
	"net/http"

	"example.com/belajar-go/config"
	"example.com/belajar-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Event{})

	server := gin.Default()

	// route
	api := server.Group("/api")
	{
		api.POST("/events", createEvents)
		api.GET("/events", getEvents)
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



func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
			"error":   err.Error(),
		})
		return
	}

	// dummy
	event.UserId = 1

	// save inputan
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"Message": "create Event",
		"event":   event,
	})
}
