package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/api"
)

func main() {
	log.Println("Creating Gin Engine...")
	router := gin.Default()
	api.SetupApi(router)
	log.Println("Running API...")
	router.Run("localhost:8080")
}
