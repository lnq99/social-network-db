package v1

import (
	"app/internal/model"
	"app/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	LoginHandler() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
	LogoutHandler() gin.HandlerFunc
	// Login(email, password string) bool
	// SignUp(email, password, name, gender, birthday string) bool
}

func (ctrl *Controller) Register(c *gin.Context) {
	var profileBody service.ProfileBody
	if err := c.ShouldBindJSON(&profileBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	err := ctrl.services.Profile.Register(profileBody)
	statusRespone(c, err)
}

func (ctrl *Controller) LoginHandler(c *gin.Context) {
	var user model.Profile
	id := 0

	token, err := c.Cookie("token")
	if err == nil {
		id, err = ctrl.auth.ParseTokenId(token)
		if err == nil && id > 0 {
			c.Set("ID", id)
			user, _ = ctrl.services.Profile.Get(id)
		}
	}

	if err != nil {
		u := service.LoginBody{}

		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}

		user, _ = ctrl.services.Profile.GetByEmail(u.Email)

		if user.Email != u.Email ||
			!ctrl.auth.ComparePassword(u.Password, user.Salt, user.Hash) {
			c.JSON(http.StatusUnauthorized, "Email or password is invalid")
			return
		}

		token, err = ctrl.auth.CreateToken(strconv.Itoa(user.Id))
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
	}

	c.SetCookie("token", token, 60*60*24, "/", ctrl.conf.Host, true, true)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func (ctrl *Controller) LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", ctrl.conf.Host, true, true)
	c.Status(http.StatusOK)
}
