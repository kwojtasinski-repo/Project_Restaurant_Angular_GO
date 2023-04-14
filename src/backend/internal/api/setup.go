package api

import (
	"log"

	"github.com/chmike/securecookie"
	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
)

func SetupApi(config config.Config, router *gin.Engine) {
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
}

func configOptions(config config.Config) {
	log.Println("Creating Auth Cookie")
	Location = config.Server.Host
	CookieHashKey = []byte(config.Server.CookieHashKey)
	var cookieErr error
	CookieIssued, cookieErr = securecookie.New(CookieSessionName, CookieHashKey, securecookie.Params{
		Path:     Location,
		Domain:   Location,
		MaxAge:   CookieLifeTime,
		HTTPOnly: true,
		Secure:   true,
	})
	if cookieErr != nil {
		log.Fatal("ERROR: Cookie cannot be issued ", cookieErr.Error())
	}
}
