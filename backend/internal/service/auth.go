package service

import (
	"app/config"
	"app/internal/middleware"
	"app/internal/model"
	"app/internal/repository"
	"app/pkg/auth"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthServiceImpl struct {
	profileRepo repository.ProfileRepo
	cookieHost  string
	manager     auth.Manager
}

type LoginObj struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type SignUpObj struct {
// 	Email    string `json:"email"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Gender   string `json:"gender"`
// 	Birthday string `json:"birthday"`
// }

func NewAuthService(repo repository.ProfileRepo, conf *config.Config) AuthService {
	return &AuthServiceImpl{repo, "localhost", auth.NewManager("id", conf.ApiSecret)}
}

func (a *AuthServiceImpl) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.Profile
		id := 0

		token, err := c.Cookie("token")
		if err == nil {
			id, err = a.manager.ParseTokenId(token)
			if err == nil && id > 0 {
				c.Set("ID", id)
				user, _ = a.profileRepo.Select(id)
			}
		}

		if err != nil {
			u := LoginObj{}

			if err := c.ShouldBindJSON(&u); err != nil {
				c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
				return
			}

			user, _ = a.profileRepo.SelectByEmail(u.Email)

			if user.Email != u.Email ||
				!a.manager.ComparePassword(u.Password, user.Salt, user.Hash) {
				c.JSON(http.StatusUnauthorized, "Email or password is invalid")
				return
			}

			token, err = a.manager.CreateToken(strconv.Itoa(user.Id))
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
		}

		c.SetCookie("token", token, 60*60*1, "/", a.cookieHost, true, true)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user":  user,
		})
	}
}

func (a *AuthServiceImpl) AuthMiddleware() gin.HandlerFunc {
	return middleware.NewAuthMiddleware(a.manager)
}

func (a *AuthServiceImpl) LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("logged out")
		c.SetCookie("token", "", -1, "/", a.cookieHost, true, true)
		c.Status(http.StatusOK)
	}
}
