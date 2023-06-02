package api

import (
	"log"
	"net/http"
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
	setupCors(router)
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
	addHealthCheck(router)
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

func addHealthCheck(router *gin.Engine) {
	router.GET("/api", healthCheck)
}

func healthCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Welcome to Restaurant API!")
}

func setupCors(router *gin.Engine) {
	log.Println("Setup Cors")
	corsConfig := CorsConfig{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Cookie"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
	}
	router.Use(CORSMiddleware(corsConfig))
}
