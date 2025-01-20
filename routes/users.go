package routes

import (
	"api/models"
	"api/utils"
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

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse request data. Maybe some required values are missing!","m2":err.Error()})
		fmt.Println(err)
		return
	}

	err = user.ValidateCredentials()

	if err != nil{
		context.JSON(http.StatusUnauthorized,gin.H{"message":err.Error()})
		fmt.Println(err)
		return
	}

	token, err := utils.GenerateToken(user.Email,user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not generate token","m2":err.Error()})
		fmt.Println(err)
		return
	}


	context.JSON(http.StatusOK,gin.H{"message":"Login Successful!","token":token})
}

func getUsers(context *gin.Context){
	users, err := models.GetAllUsers()
	if err != nil{
		
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events","m2":err.Error()})
		return
	}
	context.JSON(http.StatusOK,users)
}