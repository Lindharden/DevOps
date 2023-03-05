package main

import (
	"DevOps/globals"
	helpers "DevOps/helpers"
	routes "DevOps/routes"
)

func main() {
	gormDb := helpers.GetGormDbConnection()
	globals.SetDatabase(gormDb)
	r := routes.SetupRouter()
	r.Run(":8080")
}
