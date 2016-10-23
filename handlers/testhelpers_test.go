package handlers

import (
	"io/ioutil"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
	"github.com/tcm1911/gomediacenter"
)

///////////////
// Constants //
///////////////
const authTestHeader = `MediaBrowser Client="Emby Web Client", Device="Chrome 50.0.2661.50", DeviceId="cae2cc5be4e17f1d0a486d0c8fdb4789f4f1e99c", Version="3.0.5912.0", UserId="f40b2df070cf46e686bcbdd388d8706c"`

///////////
// Mocks //
///////////

// Setup db mock
type mockDB struct {
	mock.Mock
}

func (m *mockDB) FindItemByID(s gomediacenter.ID) (gomediacenter.MEDIATYPE, interface{}, error) {
	args := m.Called(s)
	return args.Get(0).(gomediacenter.MEDIATYPE), args.Get(1).(interface{}), args.Error(2)
}

func (m *mockDB) FindItemUserData(uid, itemID gomediacenter.ID) (*gomediacenter.ItemUserData, error) {
	args := m.Called(uid, itemID)
	return args.Get(0).(*gomediacenter.ItemUserData), args.Error(1)
}

func (m *mockDB) FindItemIntro(id gomediacenter.ID) ([]*gomediacenter.Intro, error) {
	args := m.Called(id)
	return args.Get(0).([]*gomediacenter.Intro), args.Error(1)
}

func (m *mockDB) AddNewUser(user *gomediacenter.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockDB) GetUserByID(s gomediacenter.ID) (*gomediacenter.User, error) {
	args := m.Called(s)
	return args.Get(0).(*gomediacenter.User), args.Error(1)
}

func (m *mockDB) GetAllUsers(filter map[string]interface{}) ([]*gomediacenter.User, error) {
	args := m.Called(filter)
	return args.Get(0).([]*gomediacenter.User), args.Error(1)
}

func (m *mockDB) DeleteUser(hexString gomediacenter.ID) error {
	args := m.Called(hexString)
	return args.Error(0)
}

func (m *mockDB) GetUserByName(name string) (*gomediacenter.User, error) {
	args := m.Called(name)
	return args.Get(0).(*gomediacenter.User), args.Error(1)
}

func (m *mockDB) GetUserPasswordHash(uid gomediacenter.ID) ([]byte, error) {
	args := m.Called(uid)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockDB) ChangeUserPassword(uid gomediacenter.ID, newPassword []byte) error {
	args := m.Called(uid, newPassword)
	return args.Error(0)
}

func (m *mockDB) UpdateUser(uid gomediacenter.ID, user *gomediacenter.User) error {
	args := m.Called(uid, user)
	return args.Error(0)
}

func (m *mockDB) UpdateUserPolicy(uid gomediacenter.ID, policy *gomediacenter.UserPolicy) error {
	args := m.Called(uid, policy)
	return args.Error(0)
}

func (m *mockDB) UpdateUserConfiguration(uid gomediacenter.ID, conf *gomediacenter.UserConfig) error {
	args := m.Called(uid, conf)
	return args.Error(0)
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
