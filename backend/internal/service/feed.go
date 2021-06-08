package service

import (
	"app/internal/repository"
)

type FeedServiceImpl struct {
	profile repository.ProfileRepo
	// post repository.PostRepo
	// relation repository.RelationshipRepo
}

// func NewFeedService(post repository.PostRepo, relation repository.RelationshipRepo) FeedService {
// 	return &FeedServiceImpl{post, relation}
// }

func NewFeedService(profile repository.ProfileRepo) FeedService {
	return &FeedServiceImpl{profile}
}

func (r *FeedServiceImpl) GetFeed(id, limit, offset int) (feed []int64, err error) {
	// friends, err := r.relation.Friends(id)
	// if err != nil {
	// 	return
	// }
	return r.profile.SelectFeed(id, limit, offset)
}
