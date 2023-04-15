package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserEndpoints(router *gin.RouterGroup) {
	log.Println("Setup User Endpoints")
	router.GET("/users/me", getMyProfile)
}

func getMyProfile(context *gin.Context) {
	userId := context.Keys["userId"].(int64)
	userService := createUserService()
	user, err := userService.Get(userId)

	if err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, user)
	}
	ResetObjectCreator()
}
