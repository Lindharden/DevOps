package routes

import (
	"github.com/gin-gonic/gin"

	controllers "DevOps/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/register", controllers.RegisterGetHandler())
	g.POST("/register", controllers.RegisterPostHandler())
	g.GET("/public", controllers.PublicTimelineHandler())
	g.GET("/:username", controllers.UserTimelineHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.SelfTimeline())
	g.GET("/logout", controllers.LogoutGetHandler())
	g.GET("/add_message", controllers.AddMessageHandler())
	g.GET("/:username/*action", controllers.FollowHandler())
}
