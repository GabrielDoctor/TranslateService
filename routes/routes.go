package routes

import (
	handlers "backend/handlers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupRouter() {
	// TODO: Add routes
	Router = gin.Default()
	api := Router.Group("/api")
	{
		api.POST("/translate/text", handlers.TranslateText)
		api.GET("/translate/chapter/:id", handlers.TranslateChapter)
		api.GET("/lookup", handlers.Lookup)
	}
}
