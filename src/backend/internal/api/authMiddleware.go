package api

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/settings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieValue, err := settings.CookieIssued.GetValue([]byte{}, c.Request)
		if err != nil {
			log.Println("ERROR: AuthMiddleware() ", err)
			c.AbortWithStatusJSON(401, gin.H{"errors": applicationerrors.RequiredCookie})
			return
		}
		var session dto.SessionDto
		err = json.Unmarshal(cookieValue, &session)
		if err != nil {
			writeErrorResponse(c, *applicationerrors.InternalError(err.Error()))
			return
		}

		if err := session.Validate(); err != nil {
			log.Println("ERROR: AuthMiddleware() ", err)
			c.AbortWithStatusJSON(401, gin.H{"errors": applicationerrors.InvalidCookie})
			return
		}

		sessionService, errCreateObject := createSessionService()
		if errCreateObject != nil {
			writeErrorResponse(c, *applicationerrors.InternalError(errCreateObject.Error()))
		}

		// think if cookie need to be updated after refreshed session?
		refreshedSession, errStatus := sessionService.ManageSession(session)
		if errStatus != nil {
			c.Abort()
			writeErrorResponse(c, *errStatus)
			return
		}

		addSessionProvider(refreshedSession)
		c.Keys = refreshedSession.AsMap()
		c.Next()
	}
}
