package main

import (
	"DevOps/globals"
	routes "DevOps/routes"

	"github.com/jmoiron/sqlx"
)

func main() {
	globals.DB = sqlx.MustConnect("sqlite3", globals.GetDatabasePath())
	r := routes.SetupRouter()
	r.Run(":8080")
}
