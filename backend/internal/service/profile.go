package service

import (
	"app/internal/model"
	"app/internal/repository"
	"app/pkg/auth"
)

type ProfileServiceImpl struct {
	repo repository.ProfileRepo
}

func NewProfileService(repo repository.ProfileRepo) ProfileService {
	return &ProfileServiceImpl{repo}
}

func (r *ProfileServiceImpl) Get(id int) (model.Profile, error) {
	return r.repo.Select(id)
}

func (r *ProfileServiceImpl) GetByEmail(e string) (model.Profile, error) {
	return r.repo.SelectByEmail(e)
}

func (r *ProfileServiceImpl) SearchName(id int, s string) (string, error) {
	return r.repo.SearchName(id, s)
}

func (r *ProfileServiceImpl) Register(body ProfileBody) error {
	var manager auth.Manager
	salt, hashed := manager.GetHashSalt(body.Password)
	p := model.Profile{
		Email:     body.Email,
		Name:      body.Username,
		Salt:      salt,
		Hash:      hashed,
		Gender:    body.Gender,
		Birthdate: body.Birthdate,
	}
	return r.repo.Insert(p)
}

func (r *ProfileServiceImpl) SetAvatar(p model.Photo) error {
	return r.repo.SetAvatar(p)
}

func (r *ProfileServiceImpl) ChangeInfo(id int, info InfoBody) error {
	return r.repo.ChangeInfo(id, info.Info)
}
