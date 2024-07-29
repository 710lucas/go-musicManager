package MusicPersistence

import (
	"710lucas/go-music-manager/models/music"
	"errors"
)

type MusicPersistence struct {
	Musics map[int]music.Music
}

func (mp *MusicPersistence) Init() {
	mp.Musics = make(map[int]music.Music)
}

func (musicPersistence *MusicPersistence) GenerateId() int {
	return len(musicPersistence.Musics) + 1
}

func (mp *MusicPersistence) SaveMusic(music music.Music) {
	mp.Musics[music.ID] = music
}

func (mp *MusicPersistence) GetMusicById(id int) (music.Music, error) {
	music, exists := mp.Musics[id]

	if !exists {
		return music, errors.New("music not found")
	}

	return music, nil
}

func (mp *MusicPersistence) GetAllMusics() []music.Music {
	var musics []music.Music

	for _, music := range mp.Musics {
		musics = append(musics, music)
	}

	return musics
}
