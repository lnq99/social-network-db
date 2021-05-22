package repository

import (
	"app/internal/model"
	"database/sql"
)

type Repo struct {
	Profile      ProfileRepo
	Post         PostRepo
	Comment      CommentRepo
	Reaction     ReactionRepo
	Relationship RelationshipRepo
	Notification NotificationRepo
	Album        AlbumRepo
	Photo        PhotoRepo
}

type ProfileRepo interface {
	Select(id int) (model.Profile, error)
	SelectByEmail(e string) (model.Profile, error)
	// Insert(model.Profile) error
}

type PostRepo interface {
	Select(postId int) (model.Post, error)
}

type CommentRepo interface {
	Select(postId int) ([]model.Comment, error)
}

type ReactionRepo interface {
	Select(postId int) ([]model.Reaction, error)
}

type RelationshipRepo interface {
	Select(id int) ([]model.Relationship, error)
	Friends(id int) ([]model.Relationship, error)
	Requests(id int) ([]model.Relationship, error)
}

type NotificationRepo interface {
	Select(userId int) ([]model.Notification, error)
}

type AlbumRepo interface {
	Select(userId int) ([]model.Album, error)
}

type PhotoRepo interface {
	Select(userId int) ([]model.Photo, error)
}

func NewRepo(db *sql.DB) (repo Repo) {
	repo = Repo{
		Profile:      NewProfileRepo(db),
		Post:         NewPostRepo(db),
		Comment:      NewCommentRepo(db),
		Reaction:     NewReactionRepo(db),
		Relationship: NewRelationshipRepo(db),
		Notification: NewNotificationRepo(db),
		Album:        NewAlbumRepo(db),
		Photo:        NewPhotoRepo(db),
	}
	return
}
