package middleware

import (
	"net/http"
	"open-bounties-api/services"

	"github.com/gin-gonic/gin"
)

func Admin(s *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		id, ok := userID.(float64)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		u, err := s.FindUserById(uint(id))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if u.Username != "thdelmas" && u.Username != "sebamiro" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
