package api

import "github.com/gin-gonic/gin"

func SetupApi(router *gin.Engine) {
	AddProductEndpoints(router)
	AddCategoryEndpoints(router)
}
