package controller

import (
	"app/internal/middleware"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter(c Controller) *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", c.Login())

	api := r.Group("/api", middleware.AuthMiddleware)
	{
		api.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "pong",
			})
		})
	}

	// r.StaticFS("/", http.Dir(c.Conf.StaticRoot))
	r.Use(static.Serve("/", static.LocalFile(c.Conf.StaticRoot, true)))
	r.NoRoute(func(ctx *gin.Context) {
		ctx.FileFromFS("/", http.Dir(c.Conf.StaticRoot))
	})
	return r
}
