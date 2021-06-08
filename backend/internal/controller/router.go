package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl Controller) *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", ctrl.Services.Auth.LoginHandler())

	api := r.Group("/api", ctrl.Services.Auth.AuthMiddleware())
	{
		profile := api.Group("profile")
		{
			profile.GET(":id", func(c *gin.Context) {
				id := toInt(c.Param("id"))
				ID := c.MustGet("ID")
				fmt.Println("===", ID, "===")
				fmt.Println("\n\n\nOK")
				profile, err := ctrl.Services.Profile.Get(id)
				handleRespone(c, profile, err)
			})

			profile.GET("short/:id", func(c *gin.Context) {
				id := toInt(c.Param("id"))
				profile, err := ctrl.Services.Profile.Get(id)
				if err == nil {
					c.JSON(200, gin.H{
						"id":      id,
						"name":    profile.Name,
						"avatars": profile.AvatarS,
					})
					return
				}
				handleRespone(c, profile, err)
			})
		}

		api.GET("friends", func(c *gin.Context) {
			ID := c.MustGet("ID").(int)
			friends, err := ctrl.Services.Relationship.FriendsDetail(ID)
			var s interface{}
			json.Unmarshal([]byte(friends), &s)
			handleRespone(c, s, err)
		})

		photo := api.Group("photo")
		{
			photo.GET(":id", func(c *gin.Context) {
				photo, err := ctrl.Services.Photo.Get(toInt(c.Param("id")))
				handleRespone(c, photo, err)
			})

			photo.GET("u/:id", func(c *gin.Context) {
				photos, err := ctrl.Services.Photo.GetByUserId(toInt(c.Param("id")))
				handleRespone(c, photos, err)
			})
		}

		notif := api.Group("notif")
		{
			notif.GET("", func(c *gin.Context) {
				ID := c.MustGet("ID").(int)
				notif, err := ctrl.Services.Notification.Get(ID)
				handleRespone(c, notif, err)
			})
		}

		post := api.Group("post")
		{
			post.GET(":id", func(c *gin.Context) {
				post, err := ctrl.Services.Post.Get(toInt(c.Param("id")))
				handleRespone(c, post, err)
			})

			post.GET("u/:id", func(c *gin.Context) {
				post, err := ctrl.Services.Post.GetByUserId(toInt(c.Param("id")))
				handleRespone(c, post, err)
			})
		}

		cmt := api.Group("cmt")
		{
			cmt.GET(":id", func(c *gin.Context) {
				cmt, err := ctrl.Services.Comment.GetTree(toInt(c.Param("id")))
				log.Println(cmt)
				var s interface{}
				json.Unmarshal([]byte(cmt), &s)
				handleRespone(c, s, err)
			})
		}

		api.GET("feed/:id", func(c *gin.Context) {
			// c.JSON(200, []int{1, 2, 3})
			id := toInt(c.Param("id"))
			limit := toInt(c.Query("lim"))
			offset := toInt(c.Query("off"))
			feed, err := ctrl.Services.Feed.GetFeed(id, limit, offset)
			println(limit, offset)
			handleRespone(c, feed, err)
		})

		api.GET("friend", func(c *gin.Context) {
			c.JSON(200, []int{1, 2, 3})
		})

		api.GET("logout", ctrl.Services.Auth.LogoutHandler())
	}

	// r.StaticFS("/", http.Dir(c.Conf.StaticRoot))
	r.Use(static.Serve("/", static.LocalFile(ctrl.Conf.StaticRoot, true)))
	r.NoRoute(func(ctx *gin.Context) {
		ctx.FileFromFS("/", http.Dir(ctrl.Conf.StaticRoot))
	})
	return r
}

func toInt(n string) int {
	res, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		return 0
	}
	return int(res)
}

func handleRespone(c *gin.Context, obj interface{}, serverErr error) {
	if serverErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, obj)
}
