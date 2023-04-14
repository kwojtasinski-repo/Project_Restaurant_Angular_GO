package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/api"
)

const (
	migrationsUp   = "--migrations=up"
	migrationsDown = "--migrations=down"
)

func main() {
	config := config.LoadConfig("config.yml")
	for index, cmd := range os.Args {
		if cmd == migrationsUp {
			countMigrations := ""
			if len(os.Args) > index+1 {
				countMigrations = os.Args[index+1]
			}
			runMigrations(config, countMigrations)
			return
		} else if cmd == migrationsDown {
			countMigrations := ""
			if len(os.Args) > index+1 {
				countMigrations = os.Args[index+1]
			}
			downMigrations(config, countMigrations)
			return
		}
	}

	log.Println("Creating Gin Engine...")
	router := gin.Default()
	api.SetupApi(router)
	log.Println("Running API...")
	router.Run("localhost:8080")
}
