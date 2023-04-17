package api

import (
	"log"
	"os"

	"github.com/chmike/securecookie"
	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/settings"
)

const GinMode = "GIN_MODE"

func SetupApi(config config.Config) *gin.Engine {
	log.Println("Creating Gin Engine...")
	router := gin.Default()
	gin.SetMode(os.Getenv(GinMode))
	configOptions(config)
	log.Println("Setup Endpoints")
	group := router.Group("/api")
	group.Use(AuthMiddleware())
	{
		AddProductEndpoints(group)
		AddCategoryEndpoints(group)
		AddOrderEndpoints(group)
		AddCartEndpoints(group)
		AddUserEndpoints(group)
		AddSessionEndpoints(group)
	}
	AddIdentityEndpoints(router)
	return router
}

func configOptions(config config.Config) {
	log.Println("Creating Auth Cookie")
	settings.Location = config.Server.Host
	settings.CookieHashKey = []byte(config.Server.CookieHashKey)
	var cookieErr error
	settings.CookieIssued, cookieErr = securecookie.New(settings.CookieSessionName, settings.CookieHashKey, securecookie.Params{
		Path:     settings.Location,
		Domain:   settings.Location,
		MaxAge:   settings.CookieLifeTime,
		HTTPOnly: true,
		Secure:   true,
	})
	if cookieErr != nil {
		log.Fatal("ERROR: Cookie cannot be issued ", cookieErr.Error())
	}
}
