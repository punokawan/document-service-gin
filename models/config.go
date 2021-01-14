package models

import (
	"document-service-gin/db"
	"os"
)

// Mongo Server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var Server = os.Getenv("DATABASE")

// Database name
var DatabaseName = os.Getenv("DATABASE_NAME")

// Create a connection
var DbConnect = db.NewConnection(Server, DatabaseName)