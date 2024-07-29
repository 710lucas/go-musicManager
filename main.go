package main

import (
	"710lucas/go-music-manager/models/user"
	AuthorPersistence "710lucas/go-music-manager/persistence/authorPersistence"
	MusicPersistence "710lucas/go-music-manager/persistence/musicPersistence"
	UserPersistence "710lucas/go-music-manager/persistence/userPersistence"
	AuthorService "710lucas/go-music-manager/services/authorService"
	MusicService "710lucas/go-music-manager/services/musicService"
	UserService "710lucas/go-music-manager/services/userService"
	"fmt"
)

var userId = -1
var choice = 0

var userPersistence = UserPersistence.UserPersistence{Users: make(map[int]user.User)}
var userService = UserService.UserService{}

var musicPersistence = MusicPersistence.MusicPersistence{}
var musicService = MusicService.MusicService{}

var authorPersistence = AuthorPersistence.AuthorPersistence{}
var authorService = AuthorService.AuthorService{}

func strInput(variable *string) error {
	_, err := fmt.Scanln(variable)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func intInput(variable *int) error {
	_, err := fmt.Scanln(variable)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func printMenu() {
	fmt.Println("Welcome to your song manager!")
	fmt.Println("(1) - Login as user")
	fmt.Println("(2) - Login as music author")
	fmt.Println("(-1) - Logout")
	fmt.Print("> ")
}

func userChoice() {

}

func userMenu() {

	for choice != 0 {
		fmt.Println("======== User area ========")
		if userId == -1 {
			fmt.Println("(1) - Create new account")
			fmt.Println("(2) - Login in existing account")
			fmt.Println("(0) - Go back")

			result := intInput(&choice)
			if result != nil {
				continue
			}

			if choice == 1 {
				var name string
				fmt.Println("Whats your name: ")
				strInput(&name)

				var age int
				fmt.Println("Whats your age: ")
				intInput(&age)

				user := userService.CreateNewUser(name, age)
				userId = user.ID

			} else if choice == 2 {

				fmt.Println("Choose the number of your account:")
				userService.PrintAllUsers()
				fmt.Print("> ")

				intInput(&userId)

				user, err := userPersistence.GetUserById(userId)

				if err != nil {
					choice = 0
					fmt.Println(err)
					continue
				}

				fmt.Println("Welcome, ", user.Name)

			}

		} else {

			fmt.Println("(1) - Like music")
			fmt.Println("(2) - Add favorite music")
			fmt.Println("(3) - Follow music author")
			fmt.Println("(4) - See your liked musics")
			fmt.Println("(5) - See your favorite musics")
			fmt.Println("(6) - See your authors")
			fmt.Println("(0) - Go back")

			intInput(&choice)

			switch choice {
			case 0:
				userId = -1
				break

			case 1:
				fmt.Println("Choose the number of the music you wanna like")
				musicService.PrintAllMusics()
				var musicId int
				intInput(&musicId)
				userService.AddLikedMusic(musicId, userId)

				fmt.Println("Your liked musics")

				for _, likedMusicId := range userService.GetUserById(userId).LikedMusicsId {
					music := musicService.GetMusicById(likedMusicId)
					fmt.Printf("[%d] - %v\n", music.ID, music.Name)
				}
			}

		}
	}

}

func main() {

	authorPersistence.Init()
	musicPersistence.Init()
	userPersistence.Init()

	userService.Init(
		&userPersistence,
		&musicPersistence,
		&musicService,
		&authorPersistence,
	)

	musicService.Init(
		&musicPersistence,
		&authorPersistence,
		&authorService,
	)

	authorService.Init(
		&authorPersistence,
		&musicPersistence,
		&userPersistence,
	)

	author := authorService.CreateNewAuthor("Me", 20)
	musicService.CreateNewMusic("Test", author.ID, 100, "me")

	for choice != -1 {

		printMenu()

		result := intInput(&choice)
		if result != nil {
			continue
		}

		if choice == 1 {
			userMenu()
		}

		fmt.Println()

	}

}
