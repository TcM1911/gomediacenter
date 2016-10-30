package gomediacenter

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PublicUserResponse is the struct returned when an unathorized user request a list
// of all users.
type PublicUserResponse struct {
	Name string `json:"Name"`
	ID   ID     `json:"id"`
}

// LoginRequest is the request struct for a login request.
type LoginRequest struct {
	Name     string `json:"Username"`
	Password string `json:"password"`
}

// AuthUserResponse is the struct returned to the user when a login request has be processed.
type AuthUserResponse struct {
	Token   string   `json:"AccessToken"`
	Session *Session `json:"SessionInfo"`
	User    *User    `json:"User"`
}

// NewUserRequest is the request struct for creating a new user.
type NewUserRequest struct {
	Name string `json:"Name"`
}

// PasswordRequest is the request struct for changing a users's password.
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
