package routes

import (
	"DevOps/globals"
	"DevOps/helpers"
	"DevOps/middleware"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")
	router.SetFuncMap(template.FuncMap{
		"formatDatetime": helpers.FormatDatetime,
		"gravatarUrl":    helpers.GravatarUrl,
	})
	router.LoadHTMLGlob(filepath.Join(globals.Root, "templates/*.html"))
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	simulator := router.Group("/sim")
	simulator.Use(middleware.CheckRequestFromSimulator)
	simulator.Use(middleware.SimulationRequest)
	SimulatorRoutes(simulator)

	public := router.Group("/")
	PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	PrivateRoutes(private)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

	return router
}
