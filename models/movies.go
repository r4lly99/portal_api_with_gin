package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Movie defines movie object structure
type Movie struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string        `bson:"title" json:"title"`
	Type        string        `bson:"type" json:"type"`
	Description string        `bson:"description" json:"description"`
}

var (
	//movie collection name
	moviesCollection = "movies"
)

// MovieModel defines the model structure
type MovieModel struct{}

// FindAll list of movies
func (m *MovieModel) FindAll() ([]Movie, error) {
	var movies []Movie
	// Connect to the movie collection
	collection := dbConnect.Use(databaseName, moviesCollection)
	err := collection.Find(bson.M{}).All(&movies)
	return movies, err
}

// FindByID a movie by its ID
func (m *MovieModel) FindByID(id string) (Movie, error) {
	var movie Movie
	// Connect to the movie collection
	collection := dbConnect.Use(databaseName, moviesCollection)
	err := collection.FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert a movie into database
func (m *MovieModel) Insert(movie Movie) error {
	// Connect to the movie collection
	collection := dbConnect.Use(databaseName, moviesCollection)
	err := collection.Insert(&movie)
	return err
}
