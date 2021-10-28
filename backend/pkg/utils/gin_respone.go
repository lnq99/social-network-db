package utils

import (
	"app/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonRespone(c *gin.Context, obj interface{}, serverErr error) {
	if serverErr != nil {
		logger.Err(serverErr)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(200, obj)
}

func StatusRespone(c *gin.Context, serverErr error) {
	if serverErr != nil {
		logger.Err(serverErr)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
