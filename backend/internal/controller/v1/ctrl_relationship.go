package v1

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetFriendsDetail(c *gin.Context) {
	id := toInt(c.Param("id"))
	friends, err := ctrl.services.Relationship.FriendsDetail(id)
	var s interface{}
	json.Unmarshal([]byte(friends), &s)
	jsonRespone(c, s, err)
}

func (ctrl *Controller) GetMutualFriends(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	mf, err := ctrl.services.Relationship.MutualFriends(ID, id)
	jsonRespone(c, mf, err)
}

func (ctrl *Controller) ChangeType(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	t := c.Param("type")
	err := ctrl.services.Relationship.ChangeType(ID, id, t)
	statusRespone(c, err)
}

func (ctrl *Controller) GetMutualAndType(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	t := ctrl.services.Relationship.GetRelationshipWith(ID, id)
	m, err := ctrl.services.Relationship.MutualFriends(ID, id)
	jsonRespone(c, gin.H{
		"type":   t,
		"mutual": m,
	}, err)
}
