package UserPersistence

import (
	"710lucas/go-music-manager/models/user"
	"errors"
)

type UserPersistence struct {
	Users map[int]user.User
}

func (up *UserPersistence) Init() {
	up.Users = make(map[int]user.User)
}

func (up *UserPersistence) GenerateId() int {
	return len(up.Users) + 1
}

func (up *UserPersistence) SaveUser(user user.User) {
	up.Users[user.ID] = user
}

func (up *UserPersistence) GetUserById(id int) (user.User, error) {

	user, exists := up.Users[id]

	if !exists {
		return user, errors.New("user not found")
	}

	return user, nil

}

func (up *UserPersistence) GetAllUsers() []user.User {
	var allUsers []user.User

	for _, user := range up.Users {
		allUsers = append(allUsers, user)
	}

	return allUsers
}
