package test

import (
	"app/internal/repository"
	"encoding/json"
	"fmt"
)

func RepoSelect(repo repository.Repo) {
	userId := 1
	postId := 1
	profile, _ := repo.Profile.Select(userId)
	// fmt.Println(profile)
	b, _ := json.Marshal(profile)
	println(string(b))

	post, _ := repo.Post.Select(postId)
	fmt.Println(post)

	cmts, _ := repo.Comment.Select(postId)
	for i := range cmts {
		fmt.Println(cmts[i])
	}

	reacts, _ := repo.Reaction.Select(userId)
	for i := range reacts {
		fmt.Println(reacts[i])
	}

	rels, _ := repo.Relationship.Requests(userId)
	for i := range rels {
		fmt.Println(rels[i])
	}

	notifs, _ := repo.Notification.Select(userId)
	for i := range notifs {
		fmt.Println(notifs[i])
	}

	albums, _ := repo.Album.Select(userId)
	for i := range albums {
		fmt.Println(albums[i])
	}

	photos, _ := repo.Photo.Select(userId)
	for i := range photos {
		fmt.Println(photos[i])
	}
}
