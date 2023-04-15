package api

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieValue, err := CookieIssued.GetValue([]byte{}, c.Request)
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

		sessionService := createSessionService()
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
