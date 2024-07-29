package UserService

import (
	"710lucas/go-music-manager/models/person"
	"710lucas/go-music-manager/models/user"
	AuthorPersistence "710lucas/go-music-manager/persistence/authorPersistence"
	MusicPersistence "710lucas/go-music-manager/persistence/musicPersistence"
	UserPersistence "710lucas/go-music-manager/persistence/userPersistence"
	MusicService "710lucas/go-music-manager/services/musicService"
	"fmt"
)

type UserService struct {
	userPersistence   *UserPersistence.UserPersistence
	musicPersistence  *MusicPersistence.MusicPersistence
	musicService      *MusicService.MusicService
	authorPersistence *AuthorPersistence.AuthorPersistence
}

func (us *UserService) Init(
	userPersistence *UserPersistence.UserPersistence,
	musicPersitence *MusicPersistence.MusicPersistence,
	musicService *MusicService.MusicService,
	authorPersistence *AuthorPersistence.AuthorPersistence,
) {

	us.userPersistence = userPersistence
	us.musicPersistence = musicPersitence
	us.musicService = musicService
	us.authorPersistence = authorPersistence

}

func (us *UserService) CreateNewUser(name string, age int) user.User {

	person := person.Person{
		Name: name,
		Age:  age,
	}

	user := user.User{
		ID:     us.userPersistence.GenerateId(),
		Person: person,
	}

	us.userPersistence.SaveUser(user)

	return user
}

func (us *UserService) AddFavoriteMusic(musicId int, userId int) {

	user, err := us.userPersistence.GetUserById(userId)

	if err != nil {
		fmt.Println(err)
		return
	}

	music, err := us.musicPersistence.GetMusicById(musicId)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.FavoriteMusicsId = append(user.FavoriteMusicsId, music.ID)

	us.userPersistence.SaveUser(user)

}

func (us *UserService) AddLikedMusic(musicId int, userId int) {

	user, err := us.userPersistence.GetUserById(userId)

	if err != nil {
		fmt.Println(err)
		return
	}

	music, err := us.musicPersistence.GetMusicById(musicId)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.LikedMusicsId = append(user.LikedMusicsId, music.ID)

	us.musicService.LikeMusic(music.ID)

	us.userPersistence.SaveUser(user)

}

func (us *UserService) FollowAuthor(authorId int, userId int) {

	user, err := us.userPersistence.GetUserById(userId)

	if err != nil {
		fmt.Println(err)
		return
	}

	author, err := us.authorPersistence.GetAuthorById(authorId)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.FollowingIds = append(user.FollowingIds, author.ID)
	author.FollowersIds = append(author.FollowersIds, user.ID)

	us.userPersistence.SaveUser(user)
	us.authorPersistence.SaveAuthor(author)

}

func (us *UserService) GetUserById(id int) user.User {

	user, err := us.userPersistence.GetUserById(id)

	if err != nil {
		fmt.Println(err)
		return user
	}

	return user

}

func (us *UserService) GetAllUsers() []user.User {
	return us.userPersistence.GetAllUsers()
}

func (us *UserService) PrintAllUsers() {
	for _, user := range us.GetAllUsers() {
		fmt.Printf("[%d] - %v\n", user.ID, user.Name)
	}
}
