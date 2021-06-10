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

func (r *PostServiceImpl) GetReaction(postId int) ([]int64, error) {
	return r.repo.SelectReaction(postId)
}

func (r *PostServiceImpl) Add(userId int, body model.PostBody) error {
	post := model.Post{
		UserId:   userId,
		Tags:     body.Tags,
		Content:  body.Content,
		AtchType: body.AtchType,
		AtchId:   body.AtchId,
		AtchUrl:  body.AtchUrl,
	}
	if post.AtchType == "photo" {
		photoId, err := services.Photo.UploadPhoto(model.Photo{
			UserId: post.UserId,
			Url:    post.AtchUrl,
		})

		if err != nil {
			return err
		}
		post.AtchId = int(photoId)
	}
	return r.repo.Insert(post)
}
