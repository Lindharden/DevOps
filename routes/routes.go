package routes

import (
	"github.com/gin-gonic/gin"

	controllers "DevOps/controllers"
)

func SimulatorRoutes(g *gin.RouterGroup) {
	g.GET("/latest", controllers.LatestGetHandler())
	g.GET("/msgs/:username", controllers.GetMessageUserHandler())
	g.POST("/msgs/:username", controllers.PostMessageUserHandler())
	g.GET("/msgs", controllers.GetMessageHandler())
	g.GET("/fllws/:username", controllers.SimGetFollowHandler())
	g.POST("/fllws/:username", controllers.SimFollowHandler())
	g.POST("/register", controllers.SimRegisterPostHandler())
}

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/register", controllers.RegisterGetHandler())
	g.POST("/register", controllers.RegisterPostHandler())
	g.GET("/public", controllers.PublicTimelineHandler())
	g.GET("/:username", controllers.UserTimelineHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/private", controllers.SelfTimeline())
	g.GET("/private/logout", controllers.LogoutGetHandler())
	g.POST("/private/message", controllers.AddMessageHandler())
	g.GET("/private/:username/*action", controllers.FollowHandler())
}
