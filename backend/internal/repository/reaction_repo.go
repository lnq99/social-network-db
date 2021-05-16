package repository

import (
	"app/internal/model"
	"database/sql"
)

type ReactionRepoImpl struct {
	DB *sql.DB
}

func NewReactionRepo(db *sql.DB) ReactionRepo {
	return &ReactionRepoImpl{db}
}

func (r *ReactionRepoImpl) selectById(id int, str string) (res []model.Reaction, err error) {
	rows, err := r.DB.Query(str, id)
	if err != nil {
		return
	}

	row := model.Reaction{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&row.UserId, &row.PostId, &row.T)
		if err != nil {
			panic(err)
		}
		res = append(res, row)
	}

	return res, nil
}

func (r *ReactionRepoImpl) Select(postId int) (res []model.Reaction, err error) {
	res, err = r.selectById(postId, "select * from Reaction where PostId=$1")
	return
}
