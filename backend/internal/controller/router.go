package controller

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter(c Controller) *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", c.Services.Auth.LoginHandler())

	api := r.Group("/api", c.Services.Auth.AuthMiddleware())
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
