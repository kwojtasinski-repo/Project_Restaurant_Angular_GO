package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/settings"
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
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": applicationerrors.InvalidInputData})
		return
	}

	userService, errCreateObject := CreateUserService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if session, err := userService.Login(signInDto); err != nil {
		writeErrorResponse(context, *err)
	} else {
		jsonBytes, err := json.Marshal(session)
		if err != nil {
			writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
			return
		}

		if err := settings.CookieIssued.SetValue(context.Writer, jsonBytes); err != nil {
			writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
			return
		}

		context.Writer.WriteHeader(http.StatusOK)
	}
}

func signUp(context *gin.Context) {
	var addUserDto dto.AddUserDto
	if err := context.BindJSON(&addUserDto); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": applicationerrors.InvalidInputData})
		return
	}

	userService, errCreateObject := CreateUserService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if _, err := userService.Register(&addUserDto); err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusCreated)
}

func signOut(context *gin.Context) {
	sessionId := context.Keys["sessionId"].(uuid.UUID)
	sessionService, err := createSessionService()
	if err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
		return
	}

	log.Print("User with sessionId: ", sessionId, " is trying to logout")
	if err := sessionService.RevokeSession(sessionId); err != nil {
		writeErrorResponse(context, *err)
		return
	}

	if err := settings.CookieIssued.Delete(context.Writer); err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
		return
	}

	log.Print("User with sessionId: ", sessionId, " successfully logout")
	context.Writer.WriteHeader(http.StatusOK)
}
