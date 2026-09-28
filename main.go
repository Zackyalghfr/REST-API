package main

import (
	"example.com/belajar-go/config"
	"example.com/belajar-go/controllers"
	"example.com/belajar-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Event{})

	server := gin.Default()

	api := server.Group("/api")
	{
		api.POST("/events", controllers.CreateEvents)
		api.GET("/events", controllers.GetEvents)
		api.GET("/events/:id", controllers.GetEventsById)
	}

	server.Run(":8080")
}