package schedulers

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/api"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/settings"
)

func RegisterSessionCleaner() gocron.Scheduler {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(2).Hour().Do(cleanPermanentlyExpiredSessions) // every 2 hours
	return *scheduler
}

func cleanPermanentlyExpiredSessions() {
	log.Println("Running cleanPermanentlyExpiredSessions()")
	db, err := api.CreateDatabaseConnection()
	if err != nil {
		log.Println("ERROR cleanPermanentlyExpiredSessions() ", err)
		return
	}

	sessionRepository := repositories.CreateSessionRepository(*db)
	expiryDuration := time.Duration(settings.CookieLifeTime)
	log.Println("Cleaning permanently expired sessions after ", expiryDuration)
	if err := sessionRepository.DeleteSessionsExpiredAfter(expiryDuration); err != nil {
		log.Println("ERROR cleanPermanentlyExpiredSessions() ", err)
		return
	}
}
