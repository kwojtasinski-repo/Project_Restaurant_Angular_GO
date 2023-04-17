package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func writeErrorResponse(context *gin.Context, errorStatus errors.ErrorStatus) {
	log.Println("ERROR: ", errorStatus.Status, ": ", errorStatus.Message)

	if errorStatus.Status >= 400 && errorStatus.Status <= 499 {
		if len(errorStatus.Message) > 0 {
			context.IndentedJSON(errorStatus.Status, gin.H{"errors": errorStatus.Message})
		} else {
			context.Writer.WriteHeader(errorStatus.Status)
		}
		return
	}

	if errorStatus.Status >= 500 && errorStatus.Status <= 599 {
		context.IndentedJSON(errorStatus.Status, gin.H{"errors": "Something bad happen"})
		return
	}
}
