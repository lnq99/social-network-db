package repository

import (
	"app/internal/model"
	"database/sql"
	"fmt"
)

type CommentRepoImpl struct {
	DB *sql.DB
}

func NewCommentRepo(db *sql.DB) CommentRepo {
	return &CommentRepoImpl{db}
}

func (r *CommentRepoImpl) selectById(id int, str string) (res []model.Comment, err error) {
	rows, err := r.DB.Query(str, id)
	if err != nil {
		return
	}

	row := model.Comment{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&row.Id, &row.UserId, &row.PostId, &row.ParentId, &row.Content, &row.Created)
		if err != nil {
			panic(err)
		}
		res = append(res, row)
	}

	return res, nil
}

func (r *CommentRepoImpl) Select(postId int) (res []model.Comment, err error) {
	res, err = r.selectById(postId, "select * from Comment where postId=$1")
	return
}

func (r *CommentRepoImpl) Insert(cmt model.Comment) (err error) {
	query := `insert into Comment(userId, postId, parentId, content, created)
	values ($1, $2, $3, $4, now())`
	res, err := r.DB.Exec(query, cmt.UserId, cmt.PostId, cmt.ParentId, cmt.Content)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil && count == 0 {
			err = fmt.Errorf("0 row affected")
		}
	}
	return
}
