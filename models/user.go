package models

import (
	"errors"
)

type User struct {
	ID int	`json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	passHash string
}

var users = map[int]User{
	1: {ID: 1, Username: "test1", Email: "test1@example.com", passHash: "12345"},
	2: {ID: 2, Username: "test2", Email: "test2@example.com", passHash: "98765"},
	3: {ID: 3, Username: "test3", Email: "test3@example.com", passHash: "abcdef"},
	4: {ID: 4, Username: "test4", Email: "test4@example.com", passHash: "pass23"},
	5: {ID: 5, Username: "test5", Email: "test5@example.com", passHash: "aiet38"},
}

func GetUserByID(id int) (User, error) {
	nilUser := *new(User)
	user, ok := users[id]
	if (!ok) {
		return nilUser, errors.New("unable to find user with specified ID")
	}
	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	nilUser := *new(User)
	for _, user := range users {
		if (user.Username == username) {
			return user, nil
		}
	}
	return nilUser, errors.New("unable to find user with specified Username")
}

func GetUsers() ([]User) {
	data := []User{}
	for _, user := range users {
		data = append(data, user)
	}
	return data
}

func CheckLogin(userName string, password string) (User, error) {
	user, err := GetUserByUsername(userName)
	if (err != nil || user.passHash != password) {
		return user, errors.New("unable to login")
	}
	return user, nil
}