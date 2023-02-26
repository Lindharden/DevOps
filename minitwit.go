package main

import (
	routes "DevOps/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
