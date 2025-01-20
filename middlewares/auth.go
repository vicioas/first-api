package middlewares

import (
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Unauthorized"})
		return 
	}

	uid,err := utils.VerifyToken(token)

	if err != nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Unauthorized"})
		return
	}

	context.Set("userId",uid)
	context.Next()
}