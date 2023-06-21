package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
)

func AddUserEndpoints(router *gin.RouterGroup) {
	log.Println("Setup User Endpoints")
	router.GET("/users/me", getMyProfile)
}

func getMyProfile(context *gin.Context) {
	userId := context.Keys["userId"].(dto.IdObject)
	userService, errCreateObject := createUserService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	user, err := userService.Get(userId.ValueInt)

	if err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, user)
	}
}
