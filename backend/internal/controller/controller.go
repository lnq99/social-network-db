package controller

import (
	"app/config"
	"app/internal/repository"
	"app/internal/service"
)

type Controller struct {
	// Repo     *repository.Repo
	Conf     *config.Config
	Services *service.Services
}

type Service interface {
}

func NewController(repo *repository.Repo, conf *config.Config) Controller {
	return Controller{
		// Repo:     repo,
		Conf:     conf,
		Services: service.GetServices(repo, conf),
	}
}
