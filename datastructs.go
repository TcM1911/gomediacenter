package gomediacenter

import "gopkg.in/mgo.v2/bson"

type PublicUserResponse struct {
	Name string        `json:"Name"`
	ID   bson.ObjectId `json:"id"`
}

type LoginRequest struct {
	Name     string `json:"Username"`
	Password string `json:"password"`
}

type AuthUserResponse struct {
	Token   string   `json:"AccessToken"`
	Session *Session `json:"SessionInfo"`
	User    *User    `json:"User"`
}

type NewUserRequest struct {
	Name string `json:"Name"`
}

type PasswordRequest struct {
	New     string `json:"newPassword"`
	Current string `json:"currentPassword"`
}
