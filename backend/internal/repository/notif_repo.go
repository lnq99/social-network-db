package repository

import (
	"app/internal/model"
	"database/sql"
)

type NotificationRepoImpl struct {
	DB *sql.DB
}

func NewNotificationRepo(db *sql.DB) NotificationRepo {
	return &NotificationRepoImpl{db}
}

func (r *NotificationRepoImpl) selectById(id int, str string) (res []model.Notification, err error) {
	rows, err := r.DB.Query(str, id)
	if err != nil {
		return
	}

	row := model.Notification{}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&row.Id, &row.UserId, &row.T, &row.Created, &row.FromUserId, &row.PostId, &row.CmtId)
		if err != nil {
			panic(err)
		}
		res = append(res, row)
	}

	return res, nil
}

func (r *NotificationRepoImpl) Select(postId int) (res []model.Notification, err error) {
	res, err = r.selectById(postId, "select * from Notification where PostId=$1")
	return
}
