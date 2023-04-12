package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func AddIdentityEndpoints(router *gin.Engine) {
	log.Println("Setup Identity Endpoints")
	router.POST("/api/sign-in", signIn)
	router.POST("/api/sign-up", signUp)
}

func signIn(context *gin.Context) {
	var signInDto dto.SignInDto
	if err := context.BindJSON(&signInDto); err != nil {
		return
	}

	userService := createUserService()
	if session, err := userService.Login(signInDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		jsonBytes, err := json.Marshal(session)
		if err != nil {
			writeErrorResponse(context, *errors.InternalError(err.Error()))
			return
		}

		CookieIssued.SetValue(context.Writer, jsonBytes)
	}
}

func signUp(context *gin.Context) {
	var addUserDto dto.AddUserDto
	if err := context.BindJSON(&addUserDto); err != nil {
		return
	}

	userService := createUserService()
	if _, err := userService.Register(&addUserDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.Writer.WriteHeader(http.StatusCreated)
	}
}
