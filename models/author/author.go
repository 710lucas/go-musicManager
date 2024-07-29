package author

import (
	"710lucas/go-music-manager/models/person"
)

type Author struct {
	ID int

	person.Person

	MusicsIds []int

	FollowersIds []int
}
