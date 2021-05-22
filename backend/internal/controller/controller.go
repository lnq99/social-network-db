package controller

import (
	"app/config"
	"app/internal/repository"
)

type Controller struct {
	Repo *repository.Repo
	Conf *config.Config
}
