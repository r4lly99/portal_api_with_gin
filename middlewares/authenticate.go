package middlewares

import (
	// Import Gin to abort on any errors
	"github.com/gin-gonic/gin"
	// To retrieve certain methods and User model
	"portal_api_with_gin/models"
	// Services to enable use to handle decoding jwt tokens
	"portal_api_with_gin/services"
)

// This will help in handling error response
func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

// Authenticate is a middleware that fetches user details from token
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch token from the headers
		requiredToken := c.Request.Header["Authorization"]

		// Check if the token is provided
		if len(requiredToken) == 0 {
			// Abort with error
			responseWithError(c, 403, "Please login to your account")
		}

		// Get email from encoded token
		userID, _ := services.DecodeToken(requiredToken[0])

		// Fetch user based on email retrieved from token
		result, err := new(models.UserModel).GetUserByEmail(userID)

		// Check if an account was not found
		if result.Email == "" {
			// Respond with a 404 when resource is not found
			responseWithError(c, 404, "User account not found")
			return
		}

		// Check if an error occured while fetching a user
		if err != nil {
			// Respond with an Internal Server Error
			responseWithError(c, 500, "Something went wrong giving you access")
			return
		}

		// Set the User variable so that we can easily retrieve from other middlewares
		c.Set("User", result)

		// Call the next middlware
		c.Next()
	}
}
