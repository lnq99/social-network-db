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

func (r *ReactionRepoImpl) SelectByUserPost(userId, postId int) (t string, err error) {
	row := r.DB.QueryRow("select type from Reaction where userId=$1 and postId=$2 limit 1", userId, postId)
	err = row.Scan(&t)
	if err != nil {
		err = nil
		t = ""
	}

	return
}

func (r *ReactionRepoImpl) UpdateReaction(userId, postId int, t string) (err error) {
	var query string
	var res sql.Result
	if t == "del" {
		query = `delete from Reaction
		where userId = $1 and postId = $2`
		res, err = r.DB.Exec(query, userId, postId)
	} else {
		query = `insert into Reaction values ($1, $2, $3)
		on conflict (userId, postId) do update set type = $3`
		res, err = r.DB.Exec(query, userId, postId, t)
	}
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}
