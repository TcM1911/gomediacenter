package auth

import (
	"log"
	"sync"
	"time"

	"github.com/tcm1911/gomediacenter"
)

// SimpleSessionManager implements the gomediacenter.SessionManager interface.
// It stores the sessions in memory and saves the sessions to the SessionSaver
// every minute.
type SimpleSessionManager struct {
	lock     sync.RWMutex
	sessions map[string]*gomediacenter.Session
}

// AddSession adds a new sessions to the session manager to track.
func (m *SimpleSessionManager) AddSession(session *gomediacenter.Session) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.sessions == nil {
		m.sessions = make(map[string]*gomediacenter.Session)
	}
	m.sessions[session.ID.String()] = session
}

// GetSession gets a session object from the manager.
func (m *SimpleSessionManager) GetSession(id gomediacenter.ID) *gomediacenter.Session {
	m.lock.RLock()
	defer m.lock.RUnlock()
	s, ok := m.sessions[id.String()]
	if ok {
		return s
	}
	return nil
}

// RemoveSession removes an active session so user is logged out.
func (m *SimpleSessionManager) RemoveSession(uid, sessionKey gomediacenter.ID) bool {
	log.Printf("Removing session %s for user: %s\n", sessionKey.String(), uid.String())
	// First retrieve the session and validate it.
	session := m.GetSession(sessionKey)
	if session == nil || !session.UserID.Equal(uid) {
		return false
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.sessions, session.ID.String())
	return true
}

// Run starts the session manager.
func (m *SimpleSessionManager) Run(db gomediacenter.SessionSaver) chan struct{} {
	m.lock.Lock()
	m.sessions = getSessionMapFromDB(db)
	m.lock.Unlock()

	exitChan := make(chan struct{})

	go func(abort chan struct{}) {
	loop:
		for {
			select {
			case <-time.Tick(60 * time.Second):
				m.saveSessionMap(db, m.sessions)
			case <-abort:
				m.saveSessionMap(db, m.sessions)
				break loop
			}
		}
		abort <- struct{}{}
	}(exitChan)
	return exitChan
}

func (m *SimpleSessionManager) saveSessionMap(db gomediacenter.SessionSaver, sessions map[string]*gomediacenter.Session) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	size := len(m.sessions)
	if size == 0 {
		return
	}

	a := make([]*gomediacenter.Session, size)
	i := 0
	for _, val := range m.sessions {
		if val != nil {
			a[i] = val
			i++
		}
	}

	if err := db.SaveSessions(a); err != nil {
		log.Println("Error while saving current sessions to db:", err)
	}
	log.Println(len(a), "sessions saved to the database.")
}

func getSessionMapFromDB(db gomediacenter.SessionSaver) map[string]*gomediacenter.Session {
	sessions, err := db.GetSavedSessions()
	if err != nil {
		return make(map[string]*gomediacenter.Session)
	}

	sessionsMap := make(map[string]*gomediacenter.Session)
	for _, session := range sessions {
		sessionsMap[session.ID.String()] = session
	}
	log.Println(len(sessionsMap), "sessions retrieved from the database.")
	return sessionsMap
}
