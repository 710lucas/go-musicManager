package music

import "710lucas/go-music-manager/models/author"

type Music struct {
	name     string
	author   *author.Author
	likes    int
	duration int
	lyrics   string
}
