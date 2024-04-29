package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "net/http"
    "fmt"
    "strings"
)

func AuthorizeJWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        const Bearer_schema = "Bearer "
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, Bearer_schema)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method")
            }
            return []byte("your_secret_key"), nil
        })

        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("userID", claims["user_id"])
        } else {
            c.AbortWithStatus(http.StatusUnauthorized)
        }
    }
}

