package controller

import (
	"app/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl Controller) *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", ctrl.Services.Auth.LoginHandler())

	r.POST("/api/register", func(c *gin.Context) {
		var profileBody model.ProfileBody
		if err := c.ShouldBindJSON(&profileBody); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}
		err := ctrl.Services.Profile.Register(profileBody)
		statusRespone(c, err)
	})

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
				jsonRespone(c, profile, err)
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
				jsonRespone(c, profile, err)
			})
		}

		rel := api.Group("rel")
		{
			rel.GET("friends/:id", func(c *gin.Context) {
				id := toInt(c.Param("id"))
				friends, err := ctrl.Services.Relationship.FriendsDetail(id)
				var s interface{}
				json.Unmarshal([]byte(friends), &s)
				jsonRespone(c, s, err)
			})

			rel.GET("mutual-friends", func(c *gin.Context) {
				ID := c.MustGet("ID").(int)
				id := toInt(c.Param("id"))
				mf, err := ctrl.Services.Relationship.MutualFriends(ID, id)
				jsonRespone(c, mf, err)
			})
		}

		photo := api.Group("photo")
		{
			photo.GET(":id", func(c *gin.Context) {
				photo, err := ctrl.Services.Photo.GetPhoto(toInt(c.Param("id")))
				jsonRespone(c, photo, err)
			})

			photo.GET("u/:id", func(c *gin.Context) {
				photos, err := ctrl.Services.Photo.GetPhotoByUserId(toInt(c.Param("id")))
				jsonRespone(c, photos, err)
			})
		}

		notif := api.Group("notif")
		{
			notif.GET("", func(c *gin.Context) {
				ID := c.MustGet("ID").(int)
				notif, err := ctrl.Services.Notification.Get(ID)
				jsonRespone(c, notif, err)
			})
		}

		post := api.Group("post")
		{
			post.GET(":id", func(c *gin.Context) {
				post, err := ctrl.Services.Post.Get(toInt(c.Param("id")))
				jsonRespone(c, post, err)
			})

			post.GET("u/:id", func(c *gin.Context) {
				post, err := ctrl.Services.Post.GetByUserId(toInt(c.Param("id")))
				jsonRespone(c, post, err)
			})

			post.POST("", func(c *gin.Context) {
				var postBody model.PostBody
				ID := c.MustGet("ID").(int)
				if err := c.ShouldBindJSON(&postBody); err != nil {
					c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
					return
				}
				err := ctrl.Services.Post.Add(ID, postBody)
				log.Println(err)
				statusRespone(c, err)
			})
		}

		cmt := api.Group("cmt")
		{
			cmt.GET(":id", func(c *gin.Context) {
				cmt, err := ctrl.Services.Comment.GetTree(toInt(c.Param("id")))
				var s interface{}
				json.Unmarshal([]byte(cmt), &s)
				jsonRespone(c, s, err)
			})

			cmt.POST("", func(c *gin.Context) {
				var cmtBody model.CommentBody
				ID := c.MustGet("ID").(int)
				if err := c.ShouldBindJSON(&cmtBody); err != nil {
					c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
					return
				}
				err := ctrl.Services.Comment.Add(ID, cmtBody)
				statusRespone(c, err)
			})
		}

		react := api.Group("react")
		{
			react.GET(":id", func(c *gin.Context) {
				react, err := ctrl.Services.Post.GetReaction(toInt(c.Param("id")))
				jsonRespone(c, react, err)
			})

			react.GET("u/:id", func(c *gin.Context) {
				ID := c.MustGet("ID").(int)
				react, err := ctrl.Services.Reaction.GetByUserPost(ID, toInt(c.Param("id")))
				jsonRespone(c, react, err)
			})

			react.PUT(":postId/:type", func(c *gin.Context) {
				ID := c.MustGet("ID").(int)
				postId := toInt(c.Param("postId"))
				t := c.Param("type")
				err := ctrl.Services.Reaction.UpdateReaction(ID, postId, t)
				statusRespone(c, err)
			})
		}

		api.GET("feed/:id", func(c *gin.Context) {
			// c.JSON(200, []int{1, 2, 3})
			id := toInt(c.Param("id"))
			limit := toInt(c.Query("lim"))
			offset := toInt(c.Query("off"))
			feed, err := ctrl.Services.Feed.GetFeed(id, limit, offset)
			jsonRespone(c, feed, err)
		})

		api.GET("search", func(c *gin.Context) {
			ID := c.MustGet("ID").(int)
			key := c.Query("k")
			res, err := ctrl.Services.Profile.SearchName(ID, key)
			var s interface{}
			json.Unmarshal([]byte(res), &s)
			jsonRespone(c, s, err)
		})

		api.GET("logout", ctrl.Services.Auth.LogoutHandler())
	}

	// r.StaticFS("/", http.Dir(c.Conf.StaticRoot))
	r.Use(static.Serve("/", static.LocalFile(ctrl.Conf.StaticRoot, true)))
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api") {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			c.FileFromFS("/", http.Dir(ctrl.Conf.StaticRoot))
		}
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

func jsonRespone(c *gin.Context, obj interface{}, serverErr error) {
	if serverErr != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(200, obj)
}

func statusRespone(c *gin.Context, serverErr error) {
	if serverErr != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
