package main

import (
	// Log items to the terminal
	"log"

	// Import gin for route definition
	"github.com/gin-gonic/gin"
	// Import godotenv for .env variables
	"github.com/joho/godotenv"
	// Import our app controllers
	"portal_api_with_gin/controllers"
	"portal_api_with_gin/middlewares"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
	godotenv.Load()
}

func main() {
	// Init gin router
	router := gin.Default()

	// Its great to version your API's
	v1 := router.Group("/api/v1")
	{
		// Define the user controller
		user := new(controllers.UserController)
		// Create the signup endpoint
		v1.POST("/signup", user.Signup)
		// Create the login endpoint
		v1.POST("/login", user.Login)

		movies := v1.Group("/movies")

		// Define the movies controller
		link := new(controllers.MoviesController)
		movies.Use(middlewares.Authenticate())

		{
			movies.GET("/all", link.FetchMovies)
			movies.POST("/create", link.CreateMovies)
		}
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	// Init our server
	router.Run(":5000")
}
