package auth

import (
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Manager struct {
	tokenKey      string
	signingKey    []byte
	signingMethod jwt.SigningMethod
}

func NewManager(tokenKey, signingKey string) Manager {
	rand.Seed(time.Now().UnixNano())
	return Manager{
		tokenKey:      tokenKey,
		signingKey:    []byte(signingKey),
		signingMethod: jwt.SigningMethodHS256,
	}
}
