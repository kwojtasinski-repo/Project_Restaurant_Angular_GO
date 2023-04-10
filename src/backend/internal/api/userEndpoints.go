package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
)

func AddUserEndpoints(router *gin.Engine) {
	log.Println("Setup User Endpoints")
	router.POST("/api/sign-in", signIn)
	router.POST("/api/sign-up", signUp)
}

func signIn(context *gin.Context) {
	var signInDto dto.SignInDto
	if err := context.BindJSON(&signInDto); err != nil {
		return
	}

	userService := createUserService()
	if user, err := userService.Login(signInDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, user)
	}
}

func signUp(context *gin.Context) {
	var addUserDto dto.AddUserDto
	if err := context.BindJSON(&addUserDto); err != nil {
		return
	}

	userService := createUserService()
	if user, err := userService.Register(&addUserDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusCreated, user)
	}
}
