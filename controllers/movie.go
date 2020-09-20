package controllers

import (
	"portal_api_with_gin/models"

	"github.com/gin-gonic/gin"
)

// Import movie model from the models file
var movieModel = new(models.MovieModel)

// MoviesController defines the movie controller
type MoviesController struct{}

func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

// FetchMovies controller handles fetching all movies of a specific user
func (b *MoviesController) FetchMovies(c *gin.Context) {
	user := c.MustGet("User").(models.User)

	if user.Email == "" {
		responseWithError(c, 404, "Please login")
		return
	}

	var data models.MovieModel

	results, err := data.FindAll()

	if err != nil {
		responseWithError(c, 500, "Problem fetching your movies")
		return
	}

	if results != nil {
		c.JSON(200, gin.H{"movies": results})
	} else {
		c.JSON(200, gin.H{"bookmarks": []string{}})
	}

}

// CreateMovies controller handles creating a movie of a specifi user
func (b *MoviesController) CreateMovies(c *gin.Context) {
	user := c.MustGet("User").(models.User)

	if user.Email == "" {
		responseWithError(c, 404, "Please login")
		return
	}

	var data models.Movie

	if c.BindJSON(&data) != nil {
		// specified response
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		// abort the request
		c.Abort()
		// return nothing
		return
	}

	err := movieModel.Insert(data)

	// Check if there was an error when saving user
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an movies"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "New movie registered"})

}
