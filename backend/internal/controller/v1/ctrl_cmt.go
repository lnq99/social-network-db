package v1

import (
	"app/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetTreeComment(c *gin.Context) {
	cmt, err := ctrl.services.Comment.GetTree(toInt(c.Param("id")))
	var s interface{}
	json.Unmarshal([]byte(cmt), &s)
	jsonRespone(c, s, err)
}

func (ctrl *Controller) PostComment(c *gin.Context) {
	var cmtBody service.CommentBody
	ID := c.MustGet("ID").(int)
	if err := c.ShouldBindJSON(&cmtBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	err := ctrl.services.Comment.Add(ID, cmtBody)
	statusRespone(c, err)
}
