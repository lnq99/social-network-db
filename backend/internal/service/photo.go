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

func (r *PhotoServiceImpl) GetByUserId(userId int) (res []model.Photo, err error) {
	return r.repo.SelectByUserId(userId)
}

func (r *PhotoServiceImpl) Get(id int) (model.Photo, error) {
	return r.repo.Select(id)
}
