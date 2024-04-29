package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "open-bounties-api/models"
    "open-bounties-api/services"
)

func RegisterUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
        return
    }

    user, err := services.CreateUser(newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

