package main

import (
	"DevOps/globals"
	helpers "DevOps/helpers"
	routes "DevOps/routes"

	"github.com/jmoiron/sqlx"
)

func main() {
	db := sqlx.MustConnect("sqlite3", globals.GetDatabasePath())
	gormDb := helpers.GetGormDbConnection()
	globals.SetDatabase(db, gormDb)
	r := routes.SetupRouter()
	r.Run(":8080")
}
