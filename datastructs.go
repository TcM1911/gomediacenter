package gomediacenter

import (
	"bytes"
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

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

// DecodeBody parses and decodes the response body to the interface.
func DecodeBody(r *http.Response, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}

// CreateRequestWithBody creates a new http request with the request body (v).
func CreateRequestWithBody(method, url string, v interface{}) (*http.Request, error) {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(v); err != nil {
		return nil, err
	}
	return http.NewRequest(method, url, b)
}
