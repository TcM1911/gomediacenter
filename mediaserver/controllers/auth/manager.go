package auth

import (
	"log"
	"sync"
	"time"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2"
)

var lock sync.RWMutex
var sessions map[string]*gomediacenter.Session

// AddSession adds a new sessions to the session manager to track.
func AddSession(session *gomediacenter.Session) {
	lock.Lock()
	defer lock.Unlock()
	sessions[session.ID.Hex()] = session
	go saveSessionMap(sessions)
}

// GetSession gets a session object from the manager.
func GetSession(id string) *gomediacenter.Session {
	lock.RLock()
	defer lock.RUnlock()
	s, ok := sessions[id]
	if ok {
		return s
	}
	return nil
}

// RemoveSession removes an active session so user is logged out.
func RemoveSession(uid, sessionKey string) bool {
	log.Printf("Removing session %s for user: %s\n", sessionKey, uid)
	// First retrive the session and validate it.
	session := GetSession(sessionKey)
	if session == nil || session.UserID != uid {
		return false
	}
	lock.Lock()
	defer lock.Unlock()
	delete(sessions, session.ID.Hex())
	go func(id string) {
		db := db.GetDB()
		if err := db.RemoveSession(id); err != nil {
			log.Println("Error when removing the session from the db:", err)
		}
	}(session.ID.Hex())
	return true
}

// Run starts the session manager.
func Run() chan struct{} {
	lock.Lock()
	sessions = getSessionMapFromDB()
	lock.Unlock()

	exitChan := make(chan struct{})

	go func(abort chan struct{}) {
	loop:
		for {
			select {
			case <-time.Tick(60 * time.Second):
				saveSessionMap(sessions)
			case <-abort:
				saveSessionMap(sessions)
				break loop
			}
		}
		abort <- struct{}{}
	}(exitChan)
	return exitChan
}

func saveSessionMap(sessions map[string]*gomediacenter.Session) {
	lock.RLock()
	defer lock.RUnlock()
	size := len(sessions)
	if size == 0 {
		return
	}

	a := make([]*gomediacenter.Session, size)
	i := 0
	for _, val := range sessions {
		if val != nil {
			a[i] = val
			i++
		}
	}

	db := db.GetDB()
	if err := db.SaveSessions(a); err != nil {
		log.Println("Error while saving current sessions to db:", err)
	}
	log.Println(len(a), "sessions saved to the database.")
}

func getSessionMapFromDB() map[string]*gomediacenter.Session {
	db := db.GetDB()
	sessions, err := db.GetSavedSessions()
	if err != nil {
		if err != mgo.ErrNotFound {
			log.Println("Error while loading saved sessions from the db:", err)
		}
		return make(map[string]*gomediacenter.Session)
	}

	sessionsMap := make(map[string]*gomediacenter.Session)
	for _, session := range sessions {
		sessionsMap[session.ID.Hex()] = session
	}
	log.Println(len(sessionsMap), "sessions retrieved from the database.")
	return sessionsMap
}
