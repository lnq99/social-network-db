package test

import (
	"app/internal/service"
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	// "/api/login"
	users := []service.LoginBody{
		service.LoginBody{
			Email:    "test@gmail.com",
			Password: "easypass",
		},
		service.LoginBody{
			Email:    "ok123@example.com",
			Password: "password",
		},
	}

	for i, _ := range users {
		fmt.Println(users[i])
	}
}
