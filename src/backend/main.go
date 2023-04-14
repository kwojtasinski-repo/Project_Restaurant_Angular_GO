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
	migrationsUp   = "--migrations=up"
	migrationsDown = "--migrations=down"
)

func main() {
	config := config.LoadConfig("config.yml")
	if containsString(os.Args, "migrations") {
		runMigrations(config)
	} else {
		startServer(config)
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
		if cmd == migrationsUp {
			countMigrations := ""
			if len(os.Args) > index+1 {
				countMigrations = os.Args[index+1]
			}
			migrations.RunMigrations(config, countMigrations)
			return
		} else if cmd == migrationsDown {
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
	api.SetupApi(router)
	log.Println("Running API...")
	router.Run(fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
}
