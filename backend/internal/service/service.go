package service

import (
	"app/config"
	"app/internal/model"
	"app/internal/repository"

	"github.com/gin-gonic/gin"
)

type Services struct {
	Auth         AuthService
	Profile      ProfileService
	Post         PostService
	Comment      CommentService
	Reaction     ReactionService
	Relationship RelationshipService
	Notification NotificationService
	Album        AlbumService
	Photo        PhotoService
}

type AuthService interface {
	LoginHandler() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
	// Login(email, password string) bool
	// SignUp(email, password, name, gender, birthday string) bool
}

type ProfileService interface {
	Get(id int) (model.Profile, error)
	GetByEmail(e string) (model.Profile, error)
	// Insert(model.Profile) error
}

type PostService interface {
	Get(postId int) (model.Post, error)
}

type CommentService interface {
	Get(postId int) ([]model.Comment, error)
}

type ReactionService interface {
	Get(postId int) ([]model.Reaction, error)
}

type RelationshipService interface {
	Get(id int) ([]model.Relationship, error)
	Friends(id int) ([]model.Relationship, error)
	Requests(id int) ([]model.Relationship, error)
}

type NotificationService interface {
	Get(userId int) ([]model.Notification, error)
}

type AlbumService interface {
	Get(userId int) ([]model.Album, error)
}

type PhotoService interface {
	Get(userId int) ([]model.Photo, error)
}

func NewServices(repo *repository.Repo, conf *config.Config) (services *Services) {
	services = &Services{
		Auth:         NewAuthService(repo.Profile, conf),
		Profile:      NewProfileService(repo.Profile),
		Post:         NewPostService(repo.Post),
		Comment:      NewCommentService(repo.Comment),
		Reaction:     NewReactionService(repo.Reaction),
		Relationship: NewRelationshipService(repo.Relationship),
		Notification: NewNotificationService(repo.Notification),
		Album:        NewAlbumService(repo.Album),
		Photo:        NewPhotoService(repo.Photo),
	}
	return
}
