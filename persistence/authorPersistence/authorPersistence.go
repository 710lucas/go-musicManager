package AuthorPersistence

import (
	"710lucas/go-music-manager/models/author"
	"errors"
)

type AuthorPersistence struct {
	Authors map[int]author.Author
}

func (ap *AuthorPersistence) Init() {
	ap.Authors = make(map[int]author.Author)
}

func (ap *AuthorPersistence) GenerateId() int {
	return len(ap.Authors) + 1
}

func (ap *AuthorPersistence) SaveAuthor(author author.Author) author.Author {
	ap.Authors[author.ID] = author
	return author
}

func (authorPersistence *AuthorPersistence) GetAuthorById(id int) (author.Author, error) {
	author, exists := authorPersistence.Authors[id]

	if !exists {
		return author, errors.New("author not found")
	}

	return author, nil
}

func (ap *AuthorPersistence) GetAllAuthors() []author.Author {

	var authors []author.Author

	for _, author := range ap.Authors {
		authors = append(authors, author)
	}

	return authors

}
