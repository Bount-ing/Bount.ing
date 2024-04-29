package controllers

import (
    "open-bounties-api/services"
    "open-bounties-api/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

type LoginController struct {
    userService *services.UserService
}

func NewLoginController(userService *services.UserService) *LoginController {
    return &LoginController{
        userService: userService,
    }
}

// Login function to authenticate a user
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

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

