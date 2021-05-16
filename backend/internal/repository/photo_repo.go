package repository

import (
	"app/internal/model"
	"database/sql"
)

type PhotoRepoImpl struct {
	DB *sql.DB
}

func NewPhotoRepo(db *sql.DB) PhotoRepo {
	return &PhotoRepoImpl{db}
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

func (r *PhotoRepoImpl) Select(userId int) (res []model.Photo, err error) {
	res, err = r.selectById(userId, "select * from Photo where UserId=$1")
	return
}
