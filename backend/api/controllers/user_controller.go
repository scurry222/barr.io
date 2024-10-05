// controllers/user_controller.go
package controllers

import (
	"net/http"

	"barr.io/db"
	"barr.io/graph/models"
	"github.com/gin-gonic/gin"
)

// GetUsers handles GET requests to fetch all users
func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUser handles GET requests to fetch a single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles PUT requests to update an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if result := db.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE requests to remove a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if result := db.DB.Delete(&models.User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
