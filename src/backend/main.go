package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/api"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/schedulers"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/migrations"
)

const (
	MigrationsUp   = "--migrations=up"
	MigrationsDown = "--migrations=down"
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
			migrations.UpMigrations(config, countMigrations)
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
	router := api.SetupApi(config)
	scheduler := schedulers.RegisterSessionCleaner()
	scheduler.StartAsync()
	log.Println("Running API...")
	router.Run(fmt.Sprintf(":%v", config.Server.Port))
}
