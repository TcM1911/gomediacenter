package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tcm1911/gomediacenter"
)

var db *mockDB

func init() {
	db = &mockDB{}
	db.On("GetSavedSessions").Return([]*gomediacenter.Session{&gomediacenter.Session{}}, nil)
	db.On("SaveSessions", mock.AnythingOfType("string")).Return(nil)
}

func TestSessionsManager(t *testing.T) {
	assert := assert.New(t)
	manager := SimpleSessionManager{}
	manager.Run(db)
	id := gomediacenter.NewRandomID()
	session := gomediacenter.Session{Admin: false, ID: id}

	t.Run("Adding a new session", func(t *testing.T) {
		manager.AddSession(&session)
	})

	// Parallel tests.
	t.Run("GetSession of valid session", func(t *testing.T) {
		t.Parallel()
		// Should return a session with the same data as was saved.
		assert.Equal(session, *manager.GetSession(id.String()))
	})
	t.Run("GetSession of unknown session", func(t *testing.T) {
		t.Parallel()
		// Should return a nil pointer if sessions isn't known.
		randID := gomediacenter.NewRandomID()
		assert.Nil(manager.GetSession(randID.String()))
	})
	t.Run("Remove a valid session", func(t *testing.T) {
		t.Parallel()
		id := gomediacenter.NewRandomID()
		uid := gomediacenter.NewID()
		session := &gomediacenter.Session{Admin: false, ID: id, UserID: uid.String()}
		manager.AddSession(session)
		assert.True(manager.RemoveSession(uid.String(), id.String()), "Should remove known session")
	})
	t.Run("Don't remove a invalid session", func(t *testing.T) {
		t.Parallel()
		id := gomediacenter.NewRandomID()
		uid := gomediacenter.NewID()
		invalidID := gomediacenter.NewRandomID()
		session := &gomediacenter.Session{Admin: false, ID: id, UserID: uid.String()}
		manager.AddSession(session)
		assert.False(manager.RemoveSession(uid.String(), invalidID.String()), "Should not remove an unknown session")
	})
	t.Run("Don't remove a valid session if uid is not matching", func(t *testing.T) {
		t.Parallel()
		id := gomediacenter.NewRandomID()
		uid := gomediacenter.NewID()
		invalidUID := gomediacenter.NewID()
		session := &gomediacenter.Session{Admin: false, ID: id, UserID: uid.String()}
		manager.AddSession(session)
		assert.False(manager.RemoveSession(invalidUID.String(), id.String()), "Should not remove if uid is incorrect")
	})
}

type mockDB struct {
	mock.Mock
}

func (m *mockDB) GetSavedSessions() ([]*gomediacenter.Session, error) {
	args := m.Called()
	return args.Get(0).([]*gomediacenter.Session), args.Error(1)
}

func (m *mockDB) SaveSessions(s []*gomediacenter.Session) error {
	args := m.Called(s)
	return args.Error(0)
}
