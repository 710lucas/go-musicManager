package MusicService

import (
	"710lucas/go-music-manager/models/music"
	AuthorPersistence "710lucas/go-music-manager/persistence/authorPersistence"
	MusicPersistence "710lucas/go-music-manager/persistence/musicPersistence"
	AuthorService "710lucas/go-music-manager/services/authorService"
	"fmt"
)

type MusicService struct {
	musicPersistence  *MusicPersistence.MusicPersistence
	authorPersistence *AuthorPersistence.AuthorPersistence
	authorService     *AuthorService.AuthorService
}

func (mService *MusicService) Init(

	musicPersistence *MusicPersistence.MusicPersistence,
	authorPersistence *AuthorPersistence.AuthorPersistence,
	authorService *AuthorService.AuthorService,

) {

	mService.musicPersistence = musicPersistence
	mService.authorPersistence = authorPersistence
	mService.authorService = authorService

}

func (mService *MusicService) CreateNewMusic(name string, authorId int, duration int, lyrics string) music.Music {

	author, err := mService.authorPersistence.GetAuthorById(authorId)

	if err != nil {
		fmt.Println(err)
		return music.Music{}
	}

	music := music.Music{
		Name:     name,
		Author:   &author,
		Duration: duration,
		Lyrics:   lyrics,
		ID:       mService.authorPersistence.GenerateId(),
	}

	mService.musicPersistence.SaveMusic(music)

	mService.authorService.AddMusicToAuthor(authorId, music.ID)

	return music

}

func (mService *MusicService) LikeMusic(musicId int) {

	music, err := mService.musicPersistence.GetMusicById(musicId)

	if err != nil {
		fmt.Println(err)
		return
	}

	music.Likes++

	mService.musicPersistence.SaveMusic(music)

}

func (mService *MusicService) GetAllMusics() []music.Music {
	return mService.musicPersistence.GetAllMusics()
}

func (mService *MusicService) PrintAllMusics() {
	for _, music := range mService.GetAllMusics() {
		fmt.Printf("[%d] - %v\n", music.ID, music.Name)
	}
}

func (mService *MusicService) GetMusicById(id int) music.Music {

	music, err := mService.musicPersistence.GetMusicById(id)

	if err != nil {
		fmt.Println(err)
		return music
	}

	return music

}
