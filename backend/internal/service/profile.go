package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type ProfileServiceImpl struct {
	repo repository.ProfileRepo
}

func NewProfileService(repo repository.ProfileRepo) ProfileService {
	return &ProfileServiceImpl{repo}
}

func (r *ProfileServiceImpl) Get(id int) (model.Profile, error) {
	return r.repo.Select(id)
}

func (r *ProfileServiceImpl) GetByEmail(e string) (model.Profile, error) {
	return r.repo.SelectByEmail(e)
}
