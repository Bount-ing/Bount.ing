package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidAccessToken  = errors.New("invalid access token")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrMismatchedUserID    = errors.New("refresh token and access token reference differents user id")

	ErrEmptyBearerToken = errors.New("empty bearer token")

	SigningKey = []byte(os.Getenv("JWT_SECRET"))

	accessTokenDuration = 15 * time.Minute
)

func Login(mail, password string) (models.User, error) {
	var user models.User
	if dbc := db.DB.Where("mail = ?", mail).First(&user); dbc.Error != nil {
		if dbc.Error == gorm.ErrRecordNotFound {
			return user, ErrInvalidCredentials
		}
		return user, dbc.Error
	}
	if user.Password != password {
		return user, ErrInvalidCredentials
	}

	return user, nil
}

// Generates a refresh token and stores in database
func GenerateRefreshToken(uid uint) (models.RefreshToken, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return models.RefreshToken{}, err
	}

	end := time.Now().Add(time.Hour * 4)

	tkn := models.RefreshToken{
		Value:      base64.URLEncoding.EncodeToString(tokenBytes),
		ValidUntil: end,
		UserID:     uid,
	}

	dbc := db.DB.Create(&tkn)
	if dbc.Error != nil {
		return tkn, dbc.Error
	}

	return tkn, nil
}

// Checks the user's refresh token and generates a new access token if it's valid.
func RefreshTokens(refreshToken string) (string, string, error) {
	var newRefreshTkn models.RefreshToken

	var storedToken models.RefreshToken
	dbc := db.DB.Where("value = ?", refreshToken).First(&storedToken)
	if dbc.Error != nil {
		if dbc.Error == gorm.ErrRecordNotFound {
			return "", "", ErrInvalidRefreshToken
		}
		return "", "", dbc.Error
	}

	var user models.User
	user.ID = storedToken.UserID
	dbc = db.DB.Take(&user)
	if dbc.Error != nil {
		return "", "", dbc.Error
	}

	newAccessTkn, err := GenerateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	// Delete used refresh token to invalidate it before generating a new one
	dbc = db.DB.Where("value = ?", refreshToken).Delete(&models.RefreshToken{})
	if dbc.Error != nil {
		return "", "", dbc.Error
	}

	newRefreshTkn, err = GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return newRefreshTkn.Value, newAccessTkn, nil
}

type AccessTokenClaims struct {
	UID      uint
	Username string
	Admin    bool
	jwt.RegisteredClaims
}

func GenerateAccessToken(user models.User) (string, error) {
	t := time.Now()
	end := t.Add(accessTokenDuration)

	claims := AccessTokenClaims{
		UID:      user.ID,
		Username: user.Username,
		Admin:    user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(t),
			NotBefore: jwt.NewNumericDate(t),
			ExpiresAt: jwt.NewNumericDate(end),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tknString, err := tkn.SignedString(SigningKey)

	return tknString, err
}

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		rawTkn := strings.TrimPrefix(auth, "Bearer ")
		claims := AccessTokenClaims{}
		tkn, err := jwt.ParseWithClaims(
			rawTkn,
			&claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte(SigningKey), nil
			})
		if err != nil {
			if tkn != nil && !tkn.Valid {
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "expired token"})
					return
				}
			}
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

	}
}

func AdminRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		rawTkn := strings.TrimPrefix(auth, "Bearer ")
		claims := AccessTokenClaims{}
		tkn, err := jwt.ParseWithClaims(
			rawTkn,
			&claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte(SigningKey), nil
			})
		if err != nil {
			if !tkn.Valid {
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "expired token"})
					return
				}
			}
			log.Println("failed to parse jwt:", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		} else if !claims.Admin {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"reason": "not an admin"})
			return
		}
	}
}

func GetUserFromJwt(ctx *gin.Context) (models.User, error) {
	var user models.User

	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		return user, ErrEmptyBearerToken
	}

	rawTkn := strings.TrimPrefix(auth, "Bearer ")
	claims := AccessTokenClaims{}
	_, err := jwt.ParseWithClaims(
		rawTkn,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SigningKey), nil
		})

	if err != nil {
		return user, err
	}

	dbc := db.DB.Where("id = ?", claims.UID).First(&user)
	return user, dbc.Error
}
