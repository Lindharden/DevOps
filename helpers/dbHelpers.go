package helpers

import (
	"database/sql"
	"log"
)

// Convenience method to look up the id for a username.
func get_user_id(db *sql.DB, username string) string {
	var id string
	err := db.QueryRow("select user_id from user where username = ?", username).Scan(&id)
	if err != nil { 
		log.Fatal(err)
	}
	return id
}
