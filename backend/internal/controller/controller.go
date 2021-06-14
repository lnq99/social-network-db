package controller

import (
	"app/config"
	"app/internal/repository"
	"app/internal/service"
	"app/pkg/auth"
	"app/pkg/logger"
	"os"
)

type Controller struct {
	conf     *config.Config
	services *service.Services
	auth     *auth.Manager
	logger   *logger.Logger
}

func NewController(repo *repository.Repo, conf *config.Config) *Controller {
	f, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		f = os.Stdout
	}

	return &Controller{
		conf:     conf,
		services: service.GetServices(repo, conf),
		auth:     auth.NewManager("id", conf.ApiSecret),
		logger:   logger.LoggerWithWriter(f),
	}
}
