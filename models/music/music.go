package music

import "710lucas/go-music-manager/models/author"

type Music struct {
	ID       int
	Name     string
	Author   *author.Author
	Likes    int
	Duration int
	Lyrics   string
}
