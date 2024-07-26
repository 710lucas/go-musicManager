package author

import (
	"710lucas/go-music-manager/models/person"
)

type Author struct {
	ID int

	person.Person

	musicsId []int

	followersIds []int

	likes int
}
