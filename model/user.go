package model

type User struct {
	Username string `db:"username"`
	UserId   int    `db:"user_id"`
	Email    string `db:"email"`
	PwHash   string `db:"pw_hash"`
}
