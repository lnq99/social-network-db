package controller

import (
	"app/config"
	"app/internal/repository"
	"app/internal/service"
	"app/pkg/auth"
)

type Controller struct {
	conf     *config.Config
	services *service.Services
	auth     auth.Manager
}

func NewController(repo *repository.Repo, conf *config.Config) Controller {
	return Controller{
		conf:     conf,
		services: service.GetServices(repo, conf),
		auth:     auth.NewManager("id", conf.ApiSecret),
	}
}
