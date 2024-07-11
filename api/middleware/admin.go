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
		id, ok := userID.(uint)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		u, err := s.FindUserById(id)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if u.Username != "thdelmas" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
