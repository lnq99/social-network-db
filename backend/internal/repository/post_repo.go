package repository

import (
	"app/internal/model"
	"database/sql"
	"reflect"
)

type PostRepoImpl struct {
	DB *sql.DB
}

func NewPostRepo(db *sql.DB) PostRepo {
	return &PostRepoImpl{db}
}

func (u *PostRepoImpl) Select(postId int) (post model.Post, err error) {
	row := u.DB.QueryRow("select * from Post where id=$1 limit 1", postId)

	s := reflect.ValueOf(&post).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	err = row.Scan(columns...)

	return
}
