package service

// import (
// 	"app/internal/model"
// 	"app/internal/repository"
// )

// type AlbumServiceImpl struct {
// 	repo repository.AlbumRepo
// }

// func NewAlbumService(repo repository.AlbumRepo) AlbumService {
// 	return &AlbumServiceImpl{repo}
// }

// func (r *AlbumServiceImpl) Get(userId int) (res []model.Album, err error) {
// 	return r.repo.Select(userId)
// }

// func (r *AlbumServiceImpl) GetAlbumId(userId int, album string) (int, error) {
// 	return r.repo.SelectByUserIdAndAlbumName(userId, album)
// }
