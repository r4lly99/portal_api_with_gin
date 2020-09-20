package models

import (
	"os"

	"portal_api_with_gin/db"
)

// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DATABASE_HOST")

// Database name
var databaseName = "portal_api"

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
