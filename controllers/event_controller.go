package controllers

import (
	"net/http"

	"example.com/belajar-go/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
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

// func untuk instansiasi, validasi, menyimpan ke db
func CreateEvents(context *gin.Context) {
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

// untuk mengambil events berdasarkan id
func GetEventsById(context *gin.Context) {
	paramsId := context.Param("id")

	event, err := models.GetEventById(paramsId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event tidak ditemukan!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Data tampil detail event",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {
	// ambil input: id dari URL
	paramsId := context.Param("id")

	// cari data lama, kalau tidak ada maka akan error not found 404
	event, err := models.GetEventById(paramsId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event tidak ditemukan!",
		})
		return
	}

	// timpa dengan data baru dari body
	var input models.Event
	err = context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
			"error":   err.Error(),
		})
		return
	}

	event.Name = input.Name
	event.Description = input.Description
	event.Location = input.Location
	event.Datetime = input.Datetime
	// simpan perubahan
	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event berhasil diupdate",
		"event":   event,
	})
}

// func untuk menghapus data berdasarkan id
func DeleteEvent(context *gin.Context) {
	paramsId := context.Param("id")

	event, err := models.GetEventById(paramsId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event tidak ditemukan!",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not delete event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event berhasil dihapus",
	})
}
