package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type CommentServiceImpl struct {
	repo repository.CommentRepo
}

func NewCommentService(repo repository.CommentRepo) CommentService {
	return &CommentServiceImpl{repo}
}

func (r *CommentServiceImpl) Get(postId int) (res []model.Comment, err error) {
	return r.repo.Select(postId)
}
