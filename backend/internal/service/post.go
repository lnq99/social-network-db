package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type PostServiceImpl struct {
	repo repository.PostRepo
}

func NewPostService(repo repository.PostRepo) PostService {
	return &PostServiceImpl{repo}
}

func (r *PostServiceImpl) Get(postId int) (post model.Post, err error) {
	return r.repo.Select(postId)
}

func (r *PostServiceImpl) GetByUserId(userId int) ([]int64, error) {
	return r.repo.SelectByUserId(userId)
}
