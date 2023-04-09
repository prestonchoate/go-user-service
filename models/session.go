package models

import (
	"math/rand"
)

type Session struct {
	ID int `json:"id"`
	UserId int `json:"user_id"`
	AuthKey string `json:"auth_key"`
}

var sessions = make(map[string]Session)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func GetSession(user User) (Session, error) {
	session, exists := sessions[user.Username]
	if (exists) {
		return session, nil
	}
	newSession := new(Session)
	newSession.ID = len(sessions) + 1
	newSession.UserId = user.ID
	newSession.AuthKey = newToken(10)
	sessions[user.Username] = *newSession
	return *newSession, nil
}

func newToken(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}