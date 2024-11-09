package controller

import (
	"eco-journal/config"
	"eco-journal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user entities.User

	// Bind JSON request to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, entities.BaseResponse{
			Status:  false,
			Message: "Invalid request payload",
		})
		return
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.BaseResponse{
			Status:  false,
			Message: "Error hashing password",
		})
		return
	}
	user.Password = string(hashedPassword)

	// Save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, entities.BaseResponse{
			Status:  false,
			Message: "Email already registered",
		})
		return
	}

	c.JSON(http.StatusCreated, entities.BaseResponse{
		Status:  true,
		Message: "User registered successfully",
	})
}
