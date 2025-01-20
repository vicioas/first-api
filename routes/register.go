package routes

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event ID"})
		return
	}
	event,err := models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message":"Event registered!"})
}
func cancelRegistration(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse event ID"})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not cancel registration."})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"Cancelled!"})
}