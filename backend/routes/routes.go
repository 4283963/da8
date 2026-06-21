package routes

import (
	"jade-grading/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.Static("/static", "../frontend")
	r.StaticFile("/", "../frontend/index.html")

	braceletHandler := handlers.NewBraceletHandler()

	api := r.Group("/api")
	{
		bracelets := api.Group("/bracelets")
		{
			bracelets.POST("", braceletHandler.CreateBracelet)
			bracelets.GET("", braceletHandler.GetAllBracelets)
			bracelets.GET("/:id", braceletHandler.GetBracelet)
			bracelets.DELETE("/:id", braceletHandler.DeleteBracelet)
			bracelets.GET("/:id/card", braceletHandler.DownloadCard)
		}
	}

	return r
}
