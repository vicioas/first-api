package routes

import (
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events",createEvent)
	authenticated.PUT("/events/:id",updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegistration)


	server.GET("/events",getEvents)
	server.GET("/events/:id",getEvent) // /events/1, /events/2 
	server.POST("/signup",signup)
	server.POST("/login", login)
	server.GET("/users",getUsers)
}