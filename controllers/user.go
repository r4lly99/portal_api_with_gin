package controllers

import (
	"portal_api_with_gin/forms"

	"portal_api_with_gin/models"

	"github.com/gin-gonic/gin"
)

// Import the userModel from the models
var userModel = new(models.UserModel)

// UserController defines the user controller methods
type UserController struct{}

// Signup controller handles registering a user
func (u *UserController) Signup(c *gin.Context) {
	var data forms.SignupUserCommand

	// Bind the data from the request body to the SignupUserCommand Struct
	// Also check if all fields are provided
	if c.BindJSON(&data) != nil {
		// specified response
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		// abort the request
		c.Abort()
		// return nothing
		return
	}

	/*
	   You can add your validation logic
	   here such as email

	   if regexMethodChecker(data.Email) {
	       c.JSON(400, gin.H{"message": "Email is invalid"})
	       c.Abort()
	       return
	   }
	*/

	err := userModel.Signup(data)

	// Check if there was an error when saving user
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "New user account registered"})
}
