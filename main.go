package main

import (
	"api/db"
	"api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()

	server.GET("/events",getEvents)
	server.POST("/events",createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK,events)
}

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data. Maybe some required values are missing!","m2":err.Error()})
		log.Println(err)
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated,gin.H{"message":"created","event":event})
}