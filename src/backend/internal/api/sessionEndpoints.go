package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
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
		return
	}

	sessionService, err := createSessionService()
	if err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
		return
	}

	if sessions, err := sessionService.GetUserSessions(userId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, sessions)
	}
}

func revokeAllUserSessions(context *gin.Context) {
	id := context.Param("id")
	userId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	sessionService, err := createSessionService()
	if err != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(err.Error()))
		return
	}

	if err := sessionService.RevokeAllUsersSessions(userId); err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusNoContent)
}
