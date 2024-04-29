package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "open-bounties-api/models"
    "strconv"
)

// Assuming UserController and userService are defined as shown previously

func (uc *UserController) RegisterUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    registeredUser, err := uc.userService.CreateUser(newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, registeredUser)
}

func (uc *UserController) GetUser(c *gin.Context) {
    userIdStr := c.Param("id")
    userId, err := strconv.ParseUint(userIdStr, 10, 64) // Convert to uint64

    user, err := uc.userService.FindUserById(uint(userId))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
    userIdStr := c.Param("id")
    userId, err := strconv.ParseUint(userIdStr, 10, 64) // Convert to uint64
    var updateUser models.User
    if err := c.ShouldBindJSON(&updateUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    updatedUser, err := uc.userService.UpdateUser(uint(userId), updateUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
    userIdStr := c.Param("id")
    userId, err := strconv.ParseUint(userIdStr, 10, 64) // Convert to uint64
    err = uc.userService.DeleteUser(uint(userId))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

