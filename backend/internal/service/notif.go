package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type NotificationServiceImpl struct {
	repo repository.NotificationRepo
}

func NewNotificationService(repo repository.NotificationRepo) NotificationService {
	return &NotificationServiceImpl{repo}
}

func (r *NotificationServiceImpl) Get(postId int) (res []model.Notification, err error) {
	return r.repo.Select(postId)
}
