package middleware

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		if !strings.HasPrefix(authHeader, BearerSchema) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must start with Bearer"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerSchema)

		// Parse the token with explicit validation
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify signing algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, models.ErrInvalidSigningMethod
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			// Handle specific JWT validation errors
			var message string
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				// Check if refresh token is available
				if refreshToken, err := c.Cookie("refreshTkn"); err == nil && refreshToken != "" {
					// Set header to indicate token needs refresh
					c.Header("X-Token-Expired", "true")
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "Token expired",
						"code":  "TOKEN_EXPIRED",
					})
					return
				}
				message = "Token has expired"
			case errors.Is(err, jwt.ErrTokenMalformed):
				message = "Token is malformed"
			case errors.Is(err, jwt.ErrTokenSignatureInvalid):
				message = "Token signature is invalid"
			case errors.Is(err, jwt.ErrTokenNotValidYet):
				message = "Token is not valid yet"
			default:
				message = "Invalid token"
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   message,
				"details": err.Error(),
			})
			return
		}

		// Validate claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Validate expiration
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "Token has expired",
						"code":  "TOKEN_EXPIRED",
					})
					return
				}
			}

			// Validate user ID
			if userID, ok := claims["user_id"]; ok {
				// Convert userID to uint
				var uid uint
				switch v := userID.(type) {
				case float64:
					uid = uint(v)
				case string:
					if parsed, err := strconv.ParseUint(v, 10, 32); err == nil {
						uid = uint(parsed)
					}
				}

				if uid == 0 {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
					return
				}

				// Set user ID in context
				c.Set("userID", uid)

				// Optionally check if user still exists in database
				var user models.User
				if err := db.DB.First(&user, uid).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User no longer exists"})
						return
					}
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate user"})
					return
				}

				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}
