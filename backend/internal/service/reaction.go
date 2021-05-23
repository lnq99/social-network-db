package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type ReactionServiceImpl struct {
	repo repository.ReactionRepo
}

func NewReactionService(repo repository.ReactionRepo) ReactionService {
	return &ReactionServiceImpl{repo}
}

func (r *ReactionServiceImpl) Get(postId int) (res []model.Reaction, err error) {
	return r.repo.Select(postId)
}
