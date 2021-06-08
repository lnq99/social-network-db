package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type RelationshipServiceImpl struct {
	repo repository.RelationshipRepo
}

func NewRelationshipService(repo repository.RelationshipRepo) RelationshipService {
	return &RelationshipServiceImpl{repo}
}

func (r *RelationshipServiceImpl) Get(id int) (rels []model.Relationship, err error) {
	return r.repo.Select(id)
}

func (r *RelationshipServiceImpl) Friends(id int) (rels []model.Relationship, err error) {
	return r.repo.Friends(id)
}

func (r *RelationshipServiceImpl) Requests(id int) (rels []model.Relationship, err error) {
	return r.repo.Requests(id)
}

func (r *RelationshipServiceImpl) FriendsDetail(id int) (string, error) {
	return r.repo.FriendsDetail(id)
}
