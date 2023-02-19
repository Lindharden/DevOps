package helpers

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// Convenience method to look up the id for a username.
func Get_user_id(db *sqlx.DB, username string) string {
	var id string
	err := db.QueryRow("select user_id from user where username = ?", username).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}
