package middleware

import (
	"app/pkg/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(a auth.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := 0

		token, err := c.Cookie("token")
		if err == nil {
			id, err = a.ParseTokenId(token)
			if err == nil {
				c.Set("ID", id)
				log.Println(id, token, err)
				return
			}
		}

		id, err = a.ExtractTokenID(c.Request)
		log.Println("middle", id, err)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		err = a.TokenValid(c.Request)
		log.Println(err)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// user, err := auth.ParseToken(c.Request.Context(), tokenStr)
		// if err != nil {
		// 	status := http.StatusInternalServerError
		// 	if err == auth.ErrInvalidAccessToken {
		// 		status = http.StatusUnauthorized
		// 	}

		// 	c.AbortWithStatus(status)
		// 	return
		// }

		// c.Set(auth.CtxUserKey, user)
	}
}
