package AuthorService

import (
	"710lucas/go-music-manager/models/author"
	"710lucas/go-music-manager/models/person"
	AuthorPersistence "710lucas/go-music-manager/persistence/authorPersistence"
	MusicPersistence "710lucas/go-music-manager/persistence/musicPersistence"
	UserPersistence "710lucas/go-music-manager/persistence/userPersistence"
	"fmt"
)

type AuthorService struct {
	authorPersistence *AuthorPersistence.AuthorPersistence
	musicPersistence  *MusicPersistence.MusicPersistence
	userPersistence   *UserPersistence.UserPersistence
}

func (as *AuthorService) Init(
	authorPersistence *AuthorPersistence.AuthorPersistence,
	musicPersistence *MusicPersistence.MusicPersistence,
	userPersistence *UserPersistence.UserPersistence,
) {
	as.authorPersistence = authorPersistence
	as.musicPersistence = musicPersistence
	as.userPersistence = userPersistence
}

func (authorService *AuthorService) CreateNewAuthor(name string, age int) author.Author {

	authorPerson := person.Person{
		Name: name,
		Age:  age,
	}

	newId := authorService.authorPersistence.GenerateId()

	author := author.Author{
		ID:           newId,
		Person:       authorPerson,
		MusicsIds:    []int{},
		FollowersIds: []int{},
	}

	authorService.authorPersistence.SaveAuthor(author)

	return author
}

func (authorService *AuthorService) AddMusicToAuthor(authorId int, musicId int) {

	author, err := authorService.authorPersistence.GetAuthorById(authorId)

	if err != nil {
		fmt.Println(err)
		return
	}

	music, err := authorService.musicPersistence.GetMusicById(musicId)

	if err != nil {
		fmt.Println(err)
		return
	}

	author.MusicsIds = append(author.MusicsIds, music.ID)
}

//TODO: add followers
