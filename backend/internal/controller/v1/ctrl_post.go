package v1

import (
	"app/internal/service"
	"app/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetPost(c *gin.Context) {
	post, err := ctrl.services.Post.Get(toInt(c.Param("id")))
	jsonRespone(c, post, err)
}

func (ctrl *Controller) GetPostByUserId(c *gin.Context) {
	post, err := ctrl.services.Post.GetByUserId(toInt(c.Param("id")))
	jsonRespone(c, post, err)
}

func (ctrl *Controller) PostPost(c *gin.Context) {
	var postBody service.PostBody
	ID := c.MustGet("ID").(int)
	if err := c.ShouldBindJSON(&postBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	err := ctrl.services.Post.Post(ID, postBody)
	logger.Err(err)
	statusRespone(c, err)
}

func (ctrl *Controller) DeletePost(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	err := ctrl.services.Post.Delete(ID, id)
	statusRespone(c, err)
}
