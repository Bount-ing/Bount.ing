package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrUserEmailAlreadyExist = errors.New("a user with this mail already exists")
	ErrUserAlreadyVerified   = errors.New("this user has already been activated")
	ErrUserHasNoPasswd       = errors.New("user has no password")
	ErrUserVerifCodeExpired  = errors.New("verification code is expired")
)

func CreateUser(ctx *gin.Context) {
	var user models.User

	log.Println("Creating new user")
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	log.Println(user)

	if err := controllers.CreateUser(user); err != nil {
		log.Printf("Error when creating new user:", err)
		switch err {
		case controllers.ErrUserEmailAlreadyExist:
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"reason": err.Error()})
		case controllers.ErrUserHasNoPasswd:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		}
		return
	}

	ctx.Status(http.StatusCreated)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := controllers.GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
		return
	}
	// if no users are found, return an empty array and 204 status
	if len(users) == 0 {
		ctx.JSON(http.StatusNoContent, []interface{}{})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func GetCurrentUserProfileInfo(ctx *gin.Context) {
	u, err := auth.GetUserFromJwt(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	user, err := controllers.GetUserByID(u.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	userProfile := models.UserProfileReadPayload{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Admin:    user.Admin,
	}

	identities, err := controllers.GetUserIdentities(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
		return
	}

	for _, identity := range identities {
		if identity.Host.Address == "https://stripe.com" {
			userProfile.StripeConnected = true
		} else if identity.Host.Address == "https://github.com" {
			userProfile.GithubConnected = true
		}
	}

	ctx.JSON(http.StatusOK, userProfile)
}

func ValidateUserCode(ctx *gin.Context) {
	code, ok := ctx.Params.Get("code")
	if !ok {
		log.Println(code)
		ctx.AbortWithStatusJSON(409, gin.H{"reason": "No code provided"})
		return
	}

	err := controllers.CheckUserVerifCode(code)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
		}
		return

	}

	ctx.Status(http.StatusOK)
}

func CreateUserPassword(ctx *gin.Context) {
	// get code and password from json
	var data struct {
		Code     string `json:"code"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	code := data.Code
	pwd := data.Password
	log.Println(code, pwd)

	err := controllers.CreateUserPassword(pwd, code)
	if err != nil {
		if err == controllers.ErrUserVerifCodeExpired {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"reason": err})
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
			log.Println(err)
		}
		return
	}

	ctx.Status(http.StatusOK)
}

func UpdateCurrentUserProfileInfo(ctx *gin.Context) {
	var userProfileInfoPayload models.UserProfileUpdatePayload

	u, err := auth.GetUserFromJwt(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	if err = ctx.ShouldBindJSON(&userProfileInfoPayload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err = controllers.UpdateUserProfileInfo(u.ID, userProfileInfoPayload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
		return
	}

	ctx.Status(http.StatusOK)
}
