package api

import (
	"log"

	"github.com/chmike/securecookie"
	"github.com/gin-gonic/gin"
)

func SetupApi(router *gin.Engine) {
	log.Println("Creating Auth Cookie")
	var cookieErr error
	CookieIssued, cookieErr = securecookie.New(CookieSessionName, CookieHashKey, securecookie.Params{
		Path:     Location,
		Domain:   Location,
		MaxAge:   CookieLifeTime,
		HTTPOnly: true,
		Secure:   true,
	})
	if cookieErr != nil {
		log.Fatal("ERROR: Cookie cannot be issued", cookieErr.Error())
	}
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
