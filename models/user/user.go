package user

import "710lucas/go-music-manager/models/person"

type User struct {
	ID int

	person.Person

	FavoriteMusicsId []int
	LikedMusicsId    []int
	FollowingIds     []int
}
