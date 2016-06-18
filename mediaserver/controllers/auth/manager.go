package auth

import (
	"log"

	"github.com/tcm1911/gomediacenter"
	"github.com/tcm1911/gomediacenter/db"
	"gopkg.in/mgo.v2"
)

var sessionReq chan string
var sessionRes chan *gomediacenter.Session
var newSession chan *gomediacenter.Session
var delSession chan *gomediacenter.Session

// AddSession adds a new sessions to the session manager to track.
func AddSession(session *gomediacenter.Session) {
	newSession <- session
}

// RemoveSession removes an active session so user is logged out.
func RemoveSession(uid, sessionKey string) bool {
	sessionReq <- sessionKey
	session := <-sessionRes
	if session.UserId != uid {
		return false
	}
	delSession <- session
	return true
}

// Run starts the session manager.
func Run() chan struct{} {
	sessions := getSessionMapFromDB()

	sessionReq = make(chan string)
	sessionRes = make(chan *gomediacenter.Session)
	newSession = make(chan *gomediacenter.Session)
	delSession = make(chan *gomediacenter.Session)
	exitChan := make(chan struct{})

	go func(abort chan struct{}) {
	loop:
		for {
			select {
			case req := <-sessionReq:
				sessionRes <- sessions[req]
			case session := <-newSession:
				sessions[session.Id.Hex()] = session
				saveSessionMap(sessions)
			case session := <-delSession:
				delete(sessions, session.Id.Hex())
				go func(id string) {
					db := db.GetDB()
					db.RemoveSession(id)
				}(session.Id.Hex())
			case <-abort:
				saveSessionMap(sessions)
				break loop
			}
		}
		close(sessionReq)
		close(sessionRes)
		close(newSession)
		close(delSession)
		sessions = nil
		abort <- struct{}{}
	}(exitChan)
	return exitChan
}

func saveSessionMap(sessions map[string]*gomediacenter.Session) {
	size := len(sessions)
	if size == 0 {
		return
	}

	a := make([]*gomediacenter.Session, size)
	i := 0
	for _, val := range sessions {
		if val != nil {
			a[i] = val
			i += 1
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
		sessionsMap[session.Id.Hex()] = session
	}
	log.Println(len(sessionsMap), "sessions retrieved from the database.")
	return sessionsMap
}
