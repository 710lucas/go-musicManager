package user

import "710lucas/go-music-manager/models/person"

type User struct {
	ID int

	person.Person

	favoriteMusicsId []int
	likedMusicsId    []int
	followingIds     []int
}
