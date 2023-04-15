package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/api"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/migrations"
)

const (
	MigrationsUp   = "--migrations=up"
	MigrationsDown = "--migrations=down"
	GinMode        = "GIN_MODE"
)

func main() {
	configFile := config.LoadConfig("config.yml")
	api.InitObjectCreator(configFile)
	if containsString(os.Args, "migrations") {
		runMigrations(configFile)
	} else {
		startServer(configFile)
	}
}

func containsString(strArr []string, containStr string) bool {
	for _, str := range strArr {
		if strings.Contains(str, containStr) {
			return true
		}
	}

	return false
}

func runMigrations(config config.Config) {
	for index, cmd := range os.Args {
		if cmd == MigrationsUp {
			countMigrations := ""
			if len(os.Args) > index+1 {
				countMigrations = os.Args[index+1]
			}
			migrations.RunMigrations(config, countMigrations)
			return
		} else if cmd == MigrationsDown {
			countMigrations := ""
			if len(os.Args) > index+1 {
				countMigrations = os.Args[index+1]
			}
			migrations.DownMigrations(config, countMigrations)
			return
		}
	}
}

func startServer(config config.Config) {
	log.Println("Creating Gin Engine...")
	router := gin.Default()
	gin.SetMode(os.Getenv(GinMode))
	api.SetupApi(config, router)
	log.Println("Running API...")
	router.Run(fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
}
