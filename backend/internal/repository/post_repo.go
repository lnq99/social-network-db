package repository

import (
	"app/internal/model"
	"database/sql"

	"github.com/lib/pq"
)

type PostRepoImpl struct {
	DB *sql.DB
}

func NewPostRepo(db *sql.DB) PostRepo {
	return &PostRepoImpl{db}
}

func (r *PostRepoImpl) Select(postId int) (post model.Post, err error) {
	row := r.DB.QueryRow("select * from Post where id=$1 limit 1", postId)

	// s := reflect.ValueOf(&post).Elem()
	// numCols := s.NumField()
	// columns := make([]interface{}, numCols)
	// for i := 0; i < numCols; i++ {
	// 	field := s.Field(i)
	// 	columns[i] = field.Addr().Interface()
	// }
	// err = row.Scan(columns...)

	var arr pq.Int64Array

	err = row.Scan(
		&post.Id,
		&post.UserId,
		&post.Created,
		&post.Tags,
		&post.Content,
		&post.AtchType,
		&post.AtchId,
		&post.AtchUrl,
		&arr,
		&post.CmtCount,
	)
	post.Reaction = arr

	return
}

func (r *PostRepoImpl) SelectByUserId(userId int) (posts []int64, err error) {
	row := r.DB.QueryRow("select array(select id from Post where userId=$1)", userId)

	var arr pq.Int64Array
	err = row.Scan(&arr)
	posts = arr

	return
}
