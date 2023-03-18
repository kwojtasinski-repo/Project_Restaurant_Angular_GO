package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func writeErrorResponse(context *gin.Context, errorStatus errors.ErrorStatus) {
	if errorStatus.Status >= 400 && errorStatus.Status <= 499 {
		context.IndentedJSON(errorStatus.Status, gin.H{"message": errorStatus.Message})
		return
	}

	if errorStatus.Status >= 500 && errorStatus.Status <= 599 {
		context.IndentedJSON(errorStatus.Status, gin.H{"message": "Something bad happen"})
		return
	}
}
