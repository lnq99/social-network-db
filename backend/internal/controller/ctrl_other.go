package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Feed(c *gin.Context) {
	id := toInt(c.Param("id"))
	limit := toInt(c.Query("lim"))
	offset := toInt(c.Query("off"))
	feed, err := ctrl.services.Feed.GetFeed(id, limit, offset)
	jsonRespone(c, feed, err)
}

func (ctrl *Controller) Search(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	key := c.Query("k")
	res, err := ctrl.services.Profile.SearchName(ID, key)
	var s interface{}
	json.Unmarshal([]byte(res), &s)
	jsonRespone(c, s, err)
}

func (ctrl *Controller) HandleNoRoute(c *gin.Context) {
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/api") {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.FileFromFS("/", http.Dir(ctrl.conf.StaticRoot))
	}
}
