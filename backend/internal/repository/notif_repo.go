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

func (r *NotificationRepoImpl) Select(userId int) (res []model.Notification, err error) {
	res, err = r.selectById(userId, "select * from Notification where userId=$1 order by id desc limit 20")
	return
}

func (r *NotificationRepoImpl) Insert(notif model.Notification) error {
	query := `insert into Notification(userId, type, fromUserId, postId, cmtId)
	values ($1, $2, $3, $4, $5)`
	res, err := r.DB.Exec(query, notif.UserId, notif.T, notif.FromUserId, notif.PostId, notif.CmtId)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return err
}
