package main

import (
	"DevOps/globals"
	helpers "DevOps/helpers"
	"DevOps/routes"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	gormDb := helpers.GetGormDbConnection()
	globals.SetDatabase(gormDb)
	r := routes.SetupRouter()
	r.Run(":8080")

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Println("Serving metrics on 2112")
		http.ListenAndServe(":2112", nil)
	}()
}
