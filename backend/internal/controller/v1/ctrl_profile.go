package v1

import (
	"app/internal/service"
	"app/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetProfile(c *gin.Context) {
	id := toInt(c.Param("id"))
	// ID := c.MustGet("ID")
	// fmt.Println("===", ID, "===")
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
func (ctrl *Controller) ChangeIntro(c *gin.Context) {
	var introBody service.IntroBody
	ID := c.MustGet("ID").(int)
	if err := c.ShouldBindJSON(&introBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	err := ctrl.services.Profile.ChangeIntro(ID, introBody)
	logger.Err(err)
	statusRespone(c, err)
}
