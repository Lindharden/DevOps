package main

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	middleware "DevOps/middleware"
	routes "DevOps/routes"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")
	router.SetFuncMap(template.FuncMap{
		"formatDatetime": helpers.FormatDatetime,
		"gravatarUrl":    helpers.GravatarUrl,
	})
	router.LoadHTMLGlob("templates/*.html")
	router.Use(middleware.Before())
	router.Use(middleware.After())
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	return router
}

func main() {
	r := setupRouter()
	r.Run("localhost:8081")
}
