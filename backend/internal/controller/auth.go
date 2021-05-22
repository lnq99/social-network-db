package controller

import (
	"app/pkg/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		u := auth.LoginObj{}

		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}
		log.Println(u)

		user, _ := controller.Repo.Profile.SelectByEmail(u.Email)

		log.Println(user)

		if user.Email != u.Email {
			c.JSON(http.StatusUnauthorized, "Email or password is invalid")
			return
		}
		token, err := auth.CreateToken(user.Id)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.SetCookie("token", token, 60*60*24, "/", "localhost", true, true)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user":  user,
		})
	}
}
