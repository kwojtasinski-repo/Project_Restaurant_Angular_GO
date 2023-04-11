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
			c.AbortWithStatusJSON(401, "Cookie is required")
			return
		}
		var session dto.SessionDto
		json.Unmarshal(cookieValue, &session)
		log.Println("Session: ", session)
		c.Keys = session.AsMap()
		c.Next()
	}
}
