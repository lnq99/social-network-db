package controller

import (
	"app/internal/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl *Controller) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	// r := gin.New()
	// r.Use(gin.LoggerWithWriter(io.MultiWriter(os.Stdout, f)))
	// r.Use(gin.Recovery())

	r := gin.Default()

	r.Use(middleware.LoggerMiddleware(ctrl.logger))

	r.POST("/api/login", ctrl.LoginHandler)

	r.POST("/api/register", ctrl.Register)

	api := r.Group("/api", middleware.AuthMiddleware(ctrl.auth))
	{
		profile := api.Group("profile")
		{
			profile.GET(":id", ctrl.GetProfile)
			profile.GET("short/:id", ctrl.GetShortProfile)
		}

		post := api.Group("post")
		{
			post.GET(":id", ctrl.GetPost)
			post.GET("u/:id", ctrl.GetPostByUserId)
			post.POST("", ctrl.PostPost)
			post.DELETE(":id", ctrl.DeletePost)
		}

		react := api.Group("react")
		{
			react.GET(":id", ctrl.GetReaction)
			react.GET("u/:id", ctrl.GetReactionByUserPost)
			react.PUT(":postId/:type", ctrl.PutReaction)
		}

		cmt := api.Group("cmt")
		{
			cmt.GET(":id", ctrl.GetTreeComment)
			cmt.POST("", ctrl.PostComment)
		}

		rel := api.Group("rel")
		{
			rel.GET("friends/:id", ctrl.GetFriendsDetail)
			rel.GET("mutual-friends/:id", ctrl.GetMutualFriends)
			rel.GET("mutual-type/:id", ctrl.GetMutualAndType)
			rel.PUT(":id/:type", ctrl.ChangeType)
		}

		photo := api.Group("photo")
		{
			photo.GET(":id", ctrl.GetPhoto)
			photo.GET("u/:id", ctrl.GetPhotoByUserId)
		}

		notif := api.Group("notif")
		{
			notif.GET("", ctrl.GetNotifications)
		}

		api.GET("feed/:id", ctrl.Feed)

		api.GET("search", ctrl.Search)

		api.GET("logout", ctrl.LogoutHandler)
	}

	// r.StaticFS("/", http.Dir(c.Conf.StaticRoot))
	r.Use(static.Serve("/", static.LocalFile(ctrl.conf.StaticRoot, true)))
	r.NoRoute(ctrl.HandleNoRoute)

	return r
}
