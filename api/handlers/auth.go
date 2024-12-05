package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/gin-gonic/gin"
)

func Signin(ctx *gin.Context) {
	creds := struct {
		Mail     string
		Password string
	}{}
	ctx.ShouldBindJSON(&creds)
	user, err := auth.Login(strings.ToLower(creds.Mail), creds.Password)
	if err != nil {
		switch err {
		case auth.ErrInvalidCredentials:
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "Invalid credentials"})
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
		}
		return
	}

	refreshTkn, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
		log.Println("failed to generate refresh token:", err.Error())
		return
	}

	accessTkn, err := auth.GenerateAccessToken(user)
	if err != nil {
		log.Println("failed to generate access token:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
		return
	}
	cookieName := "refreshTkn"
	maxAge := 24 * 7 * time.Hour
	domain := os.Getenv("API_DOMAIN")
	path := "/"
	secure := false
	if domain != "localhost" {
		secure = true
	}
	httpOnly := true

	ctx.SetCookie(cookieName, refreshTkn.Value, int(maxAge), path, domain, secure, httpOnly)
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  accessTkn,
		"refreshToken": refreshTkn.Value,
	})
}
