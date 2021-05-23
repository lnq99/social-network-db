package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type PhotoServiceImpl struct {
	repo repository.PhotoRepo
}

func NewPhotoService(repo repository.PhotoRepo) PhotoService {
	return &PhotoServiceImpl{repo}
}

func (r *PhotoServiceImpl) Get(userId int) (res []model.Photo, err error) {
	return r.repo.Select(userId)
}
