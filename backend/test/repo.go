package test

import (
	"app/internal/repository"
	"encoding/json"
	"fmt"
)

func RepoSelect(repo repository.Repo) {
	profile, _ := repo.Profile.Select(5)
	// fmt.Println(profile)
	b, _ := json.Marshal(profile)
	println(string(b))

	post, _ := repo.Post.Select(7)
	fmt.Println(post)

	cmts, _ := repo.Comment.Select(7)
	for i := range cmts {
		fmt.Println(cmts[i])
	}

	reacts, _ := repo.Reaction.Select(5)
	for i := range reacts {
		fmt.Println(reacts[i])
	}

	rels, _ := repo.Relationship.Requests(5)
	for i := range rels {
		fmt.Println(rels[i])
	}

	notifs, _ := repo.Notification.Select(5)
	for i := range notifs {
		fmt.Println(notifs[i])
	}

	albums, _ := repo.Album.Select(5)
	for i := range albums {
		fmt.Println(albums[i])
	}

	photos, _ := repo.Photo.Select(5)
	for i := range photos {
		fmt.Println(photos[i])
	}
}
