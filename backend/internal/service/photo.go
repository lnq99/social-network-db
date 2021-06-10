package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type PhotoServiceImpl struct {
	photo repository.PhotoRepo
	album repository.AlbumRepo
}

func NewPhotoService(photo repository.PhotoRepo, album repository.AlbumRepo) PhotoService {
	return &PhotoServiceImpl{photo, album}
}

func (r *PhotoServiceImpl) GetAlbumByUserId(userId int) (res []model.Album, err error) {
	return r.album.Select(userId)
}

func (r *PhotoServiceImpl) GetAlbumId(userId int, album string) (int, error) {
	return r.album.SelectByUserIdAndAlbumName(userId, album)
}

func (r *PhotoServiceImpl) GetPhoto(id int) (model.Photo, error) {
	return r.photo.Select(id)
}

func (r *PhotoServiceImpl) GetPhotoByUserId(userId int) (res []model.Photo, err error) {
	return r.photo.SelectByUserId(userId)
}

func (r *PhotoServiceImpl) UploadPhotoToAlbum(p model.Photo, album string) (photoId int64, err error) {
	p.AlbumId, err = r.album.SelectByUserIdAndAlbumName(p.UserId, album)
	if err != nil {
		return -1, err
	}
	return r.photo.Insert(p)
}

func (r *PhotoServiceImpl) UploadPhoto(p model.Photo) (photoId int64, err error) {
	return r.UploadPhotoToAlbum(p, "Upload")
}

func (r *PhotoServiceImpl) SetAvatar(p model.Photo) (err error) {
	_, err = r.UploadPhotoToAlbum(p, "Avatar")
	return
}
