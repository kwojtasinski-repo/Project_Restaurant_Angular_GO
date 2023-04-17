package api

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/settings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieValue, err := settings.CookieIssued.GetValue([]byte{}, c.Request)
		if err != nil {
			log.Println("ERROR: AuthMiddleware() ", err)
			c.AbortWithStatusJSON(401, gin.H{"errors": "Cookie is required"})
			return
		}
		var session dto.SessionDto
		json.Unmarshal(cookieValue, &session)
		if err := session.Validate(); err != nil {
			log.Println("ERROR: AuthMiddleware() ", err)
			c.AbortWithStatusJSON(401, gin.H{"errors": "Invalid cookie"})
			return
		}

		sessionService, errCreateObject := createSessionService()
		if errCreateObject != nil {
			writeErrorResponse(c, *applicationerrors.InternalError(errCreateObject.Error()))
		}

		refreshedSession, errStatus := sessionService.ManageSession(session)
		if errStatus != nil {
			c.Abort()
			writeErrorResponse(c, *errStatus)
			return
		}

		c.Keys = refreshedSession.AsMap()
		c.Next()
	}
}
