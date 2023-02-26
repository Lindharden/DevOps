package middleware

import (
	"DevOps/globals"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Configuration

// Make sure we are connected to the database each request and look
// up the current user so that we know he's there.
func Before() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", connectDb())
		c.Next()
	}
}

// Closes the database again at the end of the request.
func After() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		db, exists := c.Get("db")
		if exists {
			db.(*sqlx.DB).Close()
		}
	}
}

// Returns a new connection to the database
func connectDb() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", globals.DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
