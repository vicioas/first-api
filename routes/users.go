package routes

import (
	"api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data. Maybe some required values are missing!","m2":err.Error()})
		fmt.Println(err)
		return
	}
	err = user.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not save request data","m2":err.Error()})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message":"User Created Successfully"})
}