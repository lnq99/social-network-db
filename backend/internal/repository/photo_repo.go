package repository

import (
	"app/internal/model"
	"database/sql"
	"log"
)

type PhotoRepoImpl struct {
	DB *sql.DB
}

func NewPhotoRepo(db *sql.DB) PhotoRepo {
	return &PhotoRepoImpl{db}
}

func (r *PhotoRepoImpl) parsePhoto(row *sql.Row) (photo model.Photo, err error) {
	err = row.Scan(&photo.Id, &photo.UserId, &photo.AlbumId, &photo.Url, &photo.Created)

	return
}

func (r *PhotoRepoImpl) selectById(id int, str string) (res []model.Photo, err error) {
	rows, err := r.DB.Query(str, id)
	if err != nil {
		return
	}

	row := model.Photo{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&row.Id, &row.UserId, &row.AlbumId, &row.Url, &row.Created)
		if err != nil {
			panic(err)
		}
		res = append(res, row)
	}

	return res, nil
}

func (r *PhotoRepoImpl) Select(id int) (res model.Photo, err error) {
	row := r.DB.QueryRow("select * from Photo where id=$1", id)
	return r.parsePhoto(row)
}

func (r *PhotoRepoImpl) SelectByUserId(userId int) (res []model.Photo, err error) {
	res, err = r.selectById(userId, "select * from Photo where UserId=$1")
	return
}

func (r *PhotoRepoImpl) Insert(photo model.Photo) (id int64, err error) {
	query := `insert into Photo(userId, albumId, url) values ($1, $2, $3) returning id`
	err = r.DB.QueryRow(query, photo.UserId, photo.AlbumId, photo.Url).Scan(&id)
	log.Println(id, err)
	return
}
