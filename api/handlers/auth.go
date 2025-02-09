package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidAccessToken  = errors.New("invalid access token")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrMismatchedUserID    = errors.New("refresh token and access token reference differents user id")

	ErrEmptyBearerToken = errors.New("empty bearer token")

	SigningKey = []byte(os.Getenv("JWT_SECRET"))

	accessTokenDuration = 60 * 2 * time.Minute
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

func RefreshToken(ctx *gin.Context) {
	// Get refresh token from cookie
	refreshTokenValue, err := ctx.Cookie("refreshTkn")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "No refresh token provided"})
		return
	}

	// Find refresh token in database
	var refreshToken models.RefreshToken
	result := db.DB.Where("value = ?", refreshTokenValue).First(&refreshToken)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": ErrInvalidRefreshToken.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": "Database error"})
		return
	}

	// Check if refresh token is expired
	if time.Now().After(refreshToken.ValidUntil) {
		// Delete expired token
		db.DB.Delete(&refreshToken)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "Refresh token expired"})
		return
	}

	// Get user associated with refresh token
	var user models.User
	if err := db.DB.First(&user, refreshToken.UserID).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": "User not found"})
		return
	}

	// Generate new access token
	accessToken, err := auth.GenerateAccessToken(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": "Failed to generate access token"})
		return
	}

	// Generate new refresh token
	newRefreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": "Failed to generate refresh token"})
		return
	}

	// Delete old refresh token
	db.DB.Delete(&refreshToken)

	// Set new refresh token cookie
	cookieName := "refreshTkn"
	maxAge := int(24 * 7 * time.Hour.Seconds())
	domain := os.Getenv("API_DOMAIN")
	path := "/"
	secure := domain != "localhost"
	httpOnly := true

	ctx.SetCookie(cookieName, newRefreshToken.Value, maxAge, path, domain, secure, httpOnly)

	// Return new access token
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": newRefreshToken.Value,
	})
}

type ResetPasswordRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordBody struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

// RequestPasswordResetHandler handles the password reset request
func RequestPasswordReset(c *gin.Context) {
	var req ResetPasswordRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := controllers.RequestPasswordReset(req.Email)
	if err != nil {
		// Don't reveal if email exists for security
		c.JSON(http.StatusOK, gin.H{
			"message": "If your email exists in our system, you will receive reset instructions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "If your email exists in our system, you will receive reset instructions",
	})
}

// ResetPasswordHandler handles the actual password reset
func ResetPassword(c *gin.Context) {
	var req ResetPasswordBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := controllers.ResetPassword(req.Code, req.Password)
	if err != nil {
		status := http.StatusInternalServerError
		if err == controllers.ErrInvalidResetCode {
			status = http.StatusBadRequest
		} else if err == controllers.ErrResetCodeExpired {
			status = http.StatusBadRequest
		}

		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password has been reset successfully",
	})
}
