package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/settings"
)

func AddIdentityEndpoints(router *gin.Engine) {
	log.Println("Setup Identity Endpoints")
	router.POST("/api/sign-in", signIn)
	router.POST("/api/sign-up", signUp)
}

func signIn(context *gin.Context) {
	var signInDto dto.SignInDto
	if err := context.BindJSON(&signInDto); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		ResetObjectCreator()
		return
	}

	userService, errCreateObject := createUserService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if session, err := userService.Login(signInDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		jsonBytes, err := json.Marshal(session)
		if err != nil {
			writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
			ResetObjectCreator()
			return
		}

		settings.CookieIssued.SetValue(context.Writer, jsonBytes)
		ResetObjectCreator()
	}
}

func signUp(context *gin.Context) {
	var addUserDto dto.AddUserDto
	if err := context.BindJSON(&addUserDto); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		ResetObjectCreator()
		return
	}

	userService, errCreateObject := createUserService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if _, err := userService.Register(&addUserDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.Writer.WriteHeader(http.StatusCreated)
	}
	ResetObjectCreator()
}
