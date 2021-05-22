package middleware

import (
	"app/pkg/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	id := 0

	token, err := c.Cookie("token")
	if err == nil {
		id, err = auth.ParseTokenId(token)
		log.Println(id, token, err)
		return
	}

	id, err = auth.ExtractTokenID(c.Request)
	log.Println("middle", id, err)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	err = auth.TokenValid(c.Request)
	log.Println(err)

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
