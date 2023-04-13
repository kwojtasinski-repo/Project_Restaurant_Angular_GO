package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.Keys["role"].(string)

		if len(userRole) == 0 || userRole != role {
			log.Printf("ERROR: PermissionMiddleware(%v) Invalid Role \n", role)
			c.AbortWithStatus(http.StatusForbidden)
		}

		c.Next()
	}
}
