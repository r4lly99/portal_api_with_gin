package db

import (
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

// Connection defines the connection structure
type Connection struct {
	session *mgo.Session
}

// NewConnection handles connecting to a mongo database
func NewConnection(host string, dbName string) (conn *Connection) {
	info := &mgo.DialInfo{
		// Address if its a local db then the value host=localhost
		Addrs: []string{host},
		// Timeout when a failure to connect to db
		Timeout: 60 * time.Second,
		// Database name
		Database: dbName,
		// Database credentials if your db is protected
		Username: os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PWD"),
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &Connection{session}
	return conn
}

// Use handles connect to a certain collection
func (conn *Connection) Use(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *Connection) Close() {
	// This closes the connection
	conn.session.Close()
	return
}
