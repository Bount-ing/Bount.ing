package controllers

import (
    "github.com/gin-gonic/gin"
    "open-bounties-api/models"
    "open-bounties-api/services"
    "net/http"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{
        userService: userService,
    }
}

// Register function to create a new user
func (ctl *UserController) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    // Adjusting to handle two return values
    createdUser, err := ctl.userService.CreateUser(user)  // Assume CreateUser accepts a value, not a pointer
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdUser)
}

