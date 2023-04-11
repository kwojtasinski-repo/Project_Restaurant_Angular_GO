package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func AddUserEndpoints(router *gin.Engine) {
	log.Println("Setup User Endpoints")
	router.POST("/api/sign-in", signIn)
	router.POST("/api/sign-up", signUp)
	//router.GET("/api/users/me", getMyProfile)
}

func signIn(context *gin.Context) {
	var signInDto dto.SignInDto
	if err := context.BindJSON(&signInDto); err != nil {
		return
	}

	userService := createUserService()
	if user, session, err := userService.Login(signInDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		fmt.Print("TODO: Add Session Cookie")
		fmt.Println(session)
		jsonBytes, err := json.Marshal(session)
		if err != nil {
			writeErrorResponse(context, *errors.InternalError(err.Error()))
			return
		}

		CookieIssued.SetValue(context.Writer, jsonBytes)
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
