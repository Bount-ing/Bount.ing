package controllers

import (
	"net/http"
	"open-bounties-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
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
	if userIdStr == "me" {
		userId, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			return
		}

		// Make sure userId is of expected type, assuming uint64 here
		userIdUInt, ok := userId.(uint)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
			return
		}

		user, err := uc.userService.FindUserById(userIdUInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found", "details": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)

	} else {
		userId, err := strconv.ParseUint(userIdStr, 10, 64) // Convert to uint64
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format", "details": err.Error()})
			return
		}

		user, err := uc.userService.FindUserById(uint(userId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found", "details": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, _ := strconv.ParseUint(userIdStr, 10, 64) // Convert to uint64
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
	userId, _ := strconv.ParseUint(userIdStr, 10, 64) // Convert to uint64
	err := uc.userService.DeleteUser(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (ctl *UserController) ConnectStripe(c *gin.Context) {
	stripeId := c.Query("id")
	userId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unautheticated user"})
		return
	}
	if id, ok := userId.(uint); ok {
		ctl.userService.UpdateUserStripeID(id, stripeId)
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": "Unexpected userID type"})
}
