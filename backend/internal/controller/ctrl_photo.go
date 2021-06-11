package controller

import "github.com/gin-gonic/gin"

func (ctrl *Controller) GetPhoto(c *gin.Context) {
	photo, err := ctrl.services.Photo.GetPhoto(toInt(c.Param("id")))
	jsonRespone(c, photo, err)
}

func (ctrl *Controller) GetPhotoByUserId(c *gin.Context) {
	photo, err := ctrl.services.Photo.GetPhotoByUserId(toInt(c.Param("id")))
	jsonRespone(c, photo, err)
}
