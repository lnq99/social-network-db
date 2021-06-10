package service

import (
	"app/config"
	"app/internal/model"
	"app/internal/repository"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	once     sync.Once
	services *Services
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
	Feed         FeedService
}

type AuthService interface {
	LoginHandler() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
	LogoutHandler() gin.HandlerFunc
	// Login(email, password string) bool
	// SignUp(email, password, name, gender, birthday string) bool
}

type ProfileService interface {
	Get(id int) (model.Profile, error)
	GetByEmail(e string) (model.Profile, error)
	SearchName(id int, s string) (string, error)

	Register(model.ProfileBody) error
	SetAvatar(model.Photo) error
}

type PostService interface {
	Get(postId int) (model.Post, error)
	GetReaction(postId int) ([]int64, error)
	GetByUserId(userId int) ([]int64, error)
	Post(userId int, body model.PostBody) error
	Delete(userId int, postId int) error
}

type CommentService interface {
	GetTree(postId int) (string, error)
	Add(userId int, body model.CommentBody) error
}

type ReactionService interface {
	Get(postId int) ([]model.Reaction, error)
	GetByUserPost(userId, postId int) (string, error)
	UpdateReaction(userId, postId int, t string) error
}

type RelationshipService interface {
	Get(id int) ([]model.Relationship, error)
	Friends(id int) ([]model.Relationship, error)
	Requests(id int) ([]model.Relationship, error)
	FriendsDetail(id int) (string, error)
	MutualFriends(u1, u2 int) ([]int64, error)
}

type NotificationService interface {
	Get(userId int) ([]model.Notification, error)
	Add(notif model.Notification) error
}

type AlbumService interface {
}

type PhotoService interface {
	GetAlbumByUserId(userId int) ([]model.Album, error)
	GetAlbumId(userId int, album string) (int, error)

	GetPhoto(id int) (model.Photo, error)
	GetPhotoByUserId(userId int) ([]model.Photo, error)

	UploadPhotoToAlbum(p model.Photo, album string) (int64, error)
	UploadPhoto(model.Photo) (int64, error)
	SetAvatar(model.Photo) error
}

type FeedService interface {
	GetFeed(id, limit, offset int) (feed []int64, err error)
	// Get(id int, tBegin, tEnd string) (newBegin, newEnd string, posts []model.Post)
}

func GetServices(repo *repository.Repo, conf *config.Config) *Services {
	once.Do(func() {
		services = &Services{
			Auth:         NewAuthService(repo.Profile, conf),
			Profile:      NewProfileService(repo.Profile),
			Post:         NewPostService(repo.Post),
			Comment:      NewCommentService(repo.Comment),
			Reaction:     NewReactionService(repo.Reaction),
			Relationship: NewRelationshipService(repo.Relationship),
			Notification: NewNotificationService(repo.Notification),
			Photo:        NewPhotoService(repo.Photo, repo.Album),
			Feed:         NewFeedService(repo.Profile),
			// Album:        NewAlbumService(repo.Album),
		}
	})
	return services
}
