package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupApi(router *gin.Engine) {
	log.Println("Setup Endpoints")
	AddProductEndpoints(router)
	AddCategoryEndpoints(router)
	AddOrderEndpoints(router)
	AddUserEndpoints(router)
}
