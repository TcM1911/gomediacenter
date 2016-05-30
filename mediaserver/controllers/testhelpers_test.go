package controllers

import (
	"io/ioutil"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
	"github.com/tcm1911/gomediacenter"
)

///////////
// Mocks //
///////////

// Setup db mock
type mockDB struct {
	mock.Mock
}

func (m *mockDB) FindItemById(s string) (gomediacenter.MEDIATYPE, interface{}, error) {
	args := m.Called(s)
	return args.Get(0).(gomediacenter.MEDIATYPE), args.Get(1).(interface{}), args.Error(2)
}

func (m *mockDB) FindItemUserData(uid, itemId string) (*gomediacenter.ItemUserData, error) {
	args := m.Called(uid, itemId)
	return args.Get(0).(*gomediacenter.ItemUserData), args.Error(1)
}

func (m *mockDB) FindItemIntro(id string) ([]*gomediacenter.Intro, error) {
	args := m.Called(id)
	return args.Get(0).([]*gomediacenter.Intro), args.Error(1)
}

func (m *mockDB) AddNewUser(user *gomediacenter.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockDB) GetUserById(s string) (*gomediacenter.User, error) {
	args := m.Called(s)
	return args.Get(0).(*gomediacenter.User), args.Error(1)
}

func (m *mockDB) GetAllUsers(filter map[string]interface{}) ([]*gomediacenter.User, error) {
	args := m.Called(filter)
	return args.Get(0).([]*gomediacenter.User), args.Error(1)
}

//////////////////////
// Helper functions //
//////////////////////

func getBodyStringFromRecorder(recorder *httptest.ResponseRecorder) (string, error) {
	p, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		return "", err
	}
	return string(p), nil
}
