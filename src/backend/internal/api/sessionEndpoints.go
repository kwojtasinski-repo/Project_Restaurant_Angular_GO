package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func AddSessionEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Session Endpoints")
	adminRouter := router.Group("/sessions", PermissionMiddleware("admin"))
	adminRouter.GET("/:id", getAllUserSessions)
	adminRouter.DELETE("/:id", revokeAllUserSessions)
}

func getAllUserSessions(context *gin.Context) {
	id := context.Param("id")
	userId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	sessionService, err := createSessionService()
	if err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
	}

	if sessions, err := sessionService.GetUserSessions(userId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, sessions)
	}
	ResetObjectCreator()
}

func revokeAllUserSessions(context *gin.Context) {
	id := context.Param("id")
	userId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	sessionService, err := createSessionService()
	if err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
	}

	if err := sessionService.RevokeAllUsersSessions(userId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.Writer.WriteHeader(http.StatusNoContent)
	}
	ResetObjectCreator()
}
