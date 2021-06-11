package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetProfile(c *gin.Context) {
	id := toInt(c.Param("id"))
	ID := c.MustGet("ID")
	fmt.Println("===", ID, "===")
	fmt.Println("\n\n\nOK")
	profile, err := ctrl.services.Profile.Get(id)
	jsonRespone(c, profile, err)
}

func (ctrl *Controller) GetShortProfile(c *gin.Context) {
	id := toInt(c.Param("id"))
	profile, err := ctrl.services.Profile.Get(id)
	if err == nil {
		c.JSON(200, gin.H{
			"id":      id,
			"name":    profile.Name,
			"avatars": profile.AvatarS,
		})
		return
	}
	jsonRespone(c, profile, err)
}
