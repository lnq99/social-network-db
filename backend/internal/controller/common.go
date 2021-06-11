package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func toInt(n string) int {
	res, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		return 0
	}
	return int(res)
}

func jsonRespone(c *gin.Context, obj interface{}, serverErr error) {
	if serverErr != nil {
		log.Println(serverErr)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(200, obj)
}

func statusRespone(c *gin.Context, serverErr error) {
	if serverErr != nil {
		log.Println(serverErr)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
