package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "open-bounties-api/models"
    "open-bounties-api/services"
)

func Login(c *gin.Context) {
    var loginParams models.LoginRequest
    if err := c.ShouldBindJSON(&loginParams); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
        return
    }

    user, token, err := services.AuthenticateUser(loginParams)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user, "token": token})
}

