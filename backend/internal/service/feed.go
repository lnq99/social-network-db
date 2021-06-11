package service

import (
	"app/internal/repository"
)

type FeedServiceImpl struct {
	profile repository.ProfileRepo
}

func NewFeedService(profile repository.ProfileRepo) FeedService {
	return &FeedServiceImpl{profile}
}

func (r *FeedServiceImpl) GetFeed(id, limit, offset int) (feed []int64, err error) {
	return r.profile.SelectFeed(id, limit, offset)
}
