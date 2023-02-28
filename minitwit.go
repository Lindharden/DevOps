package main

import (
	"DevOps/globals"
	routes "DevOps/routes"

	"github.com/jmoiron/sqlx"
)

func main() {
	db := sqlx.MustConnect("sqlite3", globals.GetDatabasePath())
	globals.SetDatabase(db)
	r := routes.SetupRouter()
	r.Run(":8080")
}
