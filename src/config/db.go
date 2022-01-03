package config

import (
	"log"
	"os"

	controllers "github.com/airaider/go_db/controllers"
	"github.com/go-pg/pg/v9"
)

// Connect Connection to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "testuser",
		Password: "testpass",
		Addr:     "localhost:5432",
		Database: "study",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
