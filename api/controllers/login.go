package controllers

import (
    "open-bounties-api/services"
    "open-bounties-api/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

type LoginController struct {
    userService *services.UserService
}

func NewLoginController(userService *services.UserService) *LoginController {
    return &LoginController{
        userService: userService,
    }
}
func (ctl *LoginController) Login(c *gin.Context) {
    var loginReq models.LoginRequest
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    user, err := ctl.userService.AuthenticateUser(loginReq.Username, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    // Create the JWT token
    claims := jwt.MapClaims{
        "user_id": user.ID,  // Make sure this claim is expected in the middleware
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString([]byte("your_secret_key")) // Same key as in middleware
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString, "user": user})
}