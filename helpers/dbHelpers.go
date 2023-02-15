package helpers

import (
	"github.com/jmoiron/sqlx"
)

// Convenience method to look up the id for a username.
func GetUserId(db *sqlx.DB, username string) (string, error) {
	var id string
	err := db.QueryRow("select user_id from user where username = ?", username).Scan(&id)
	return id, err
}
