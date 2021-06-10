package repository

import (
	"app/internal/model"
	"database/sql"
)

type AlbumRepoImpl struct {
	DB *sql.DB
}

func NewAlbumRepo(db *sql.DB) AlbumRepo {
	return &AlbumRepoImpl{db}
}

func (r *AlbumRepoImpl) selectById(id int, str string) (res []model.Album, err error) {
	rows, err := r.DB.Query(str, id)
	if err != nil {
		return
	}

	row := model.Album{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&row.Id, &row.UserId, &row.Descr, &row.Created)
		if err != nil {
			panic(err)
		}
		res = append(res, row)
	}

	return res, nil
}

func (r *AlbumRepoImpl) Select(userId int) (res []model.Album, err error) {
	res, err = r.selectById(userId, "select * from Album where UserId=$1")
	return
}

func (r *AlbumRepoImpl) SelectByUserIdAndAlbumName(userId int, album string) (id int, err error) {
	row := r.DB.QueryRow("select id from Album where userId=$1 and descr=$2", userId, album)
	err = row.Scan(&id)
	return
}
