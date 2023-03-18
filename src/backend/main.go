package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/api"
)

func main() {
	fmt.Println("Hello, World!")

	// TODO Add logs
	router := gin.Default()
	api.SetupApi(router)
	router.Run("localhost:8080")
}
