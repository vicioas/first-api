package routes

import (
	"api/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil{
		
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events","m2":err.Error()})
		return
	}
	context.JSON(http.StatusOK,events)
}	

func getEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event ID"})
		return
	}
	event,err := models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
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

func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event ID"})
		return
	}
	_,err = models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data. Maybe some required values are missing!","m2":err.Error()})
		log.Println(err)
		return
	}
	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data. Maybe some required values are missing!","m2":err.Error()})
		log.Println(err)
		return
	}
	updatedEvent.ID = eventId
	fmt.Println(updatedEvent)
	err = updatedEvent.Update()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not update event","m2":err.Error()})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event ID"})
		return
	}
	event ,err := models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data. Maybe some required values are missing!","m2":err.Error()})
		log.Println(err)
		return
	}
	err = event.Delete()

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not delete event","m2":err.Error()})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message": "Event deleted successfully!"})
}