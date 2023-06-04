package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/settings"
)

func AddIdentityEndpoints(router *gin.Engine) {
	log.Println("Setup Identity Endpoints")
	router.POST("/api/sign-in", signIn)
	router.POST("/api/sign-up", signUp)
	authorizedIdentityEndpoints := router.Group("/api", AuthMiddleware())
	authorizedIdentityEndpoints.POST("/sign-out", signOut)
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
		ResetObjectCreator()
		return
	}

	if session, err := userService.Login(signInDto); err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
	} else {
		jsonBytes, err := json.Marshal(session)
		if err != nil {
			writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
			ResetObjectCreator()
			return
		}

		if err := settings.CookieIssued.SetValue(context.Writer, jsonBytes); err != nil {
			writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
			ResetObjectCreator()
			return
		}

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

func signOut(context *gin.Context) {
	sessionId := context.Keys["sessionId"].(uuid.UUID)
	sessionService, err := createSessionService()
	if err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
		ResetObjectCreator()
		return
	}

	log.Print("User with sessionId: ", sessionId, " is trying to logout")
	if err := sessionService.RevokeSession(sessionId); err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
		return
	}

	if err := settings.CookieIssued.Delete(context.Writer); err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
		ResetObjectCreator()
		return
	}

	log.Print("User with sessionId: ", sessionId, " successfully logout")
	context.Writer.WriteHeader(http.StatusOK)
	ResetObjectCreator()
}
