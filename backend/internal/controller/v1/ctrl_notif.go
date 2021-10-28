package v1

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetNotifications(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	notif, err := ctrl.services.Notification.Get(ID)
	jsonRespone(c, notif, err)
}
