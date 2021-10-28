package v1

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetReaction(c *gin.Context) {
	react, err := ctrl.services.Post.GetReaction(toInt(c.Param("id")))
	jsonRespone(c, react, err)
}

func (ctrl *Controller) GetReactionByUserPost(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	react, err := ctrl.services.Reaction.GetByUserPost(ID, toInt(c.Param("id")))
	jsonRespone(c, react, err)
}

func (ctrl *Controller) PutReaction(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	postId := toInt(c.Param("postId"))
	t := c.Param("type")
	err := ctrl.services.Reaction.UpdateReaction(ID, postId, t)
	statusRespone(c, err)
}
