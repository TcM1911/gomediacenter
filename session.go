package gomediacenter

// Session holds all the data for an authenticated session.
type Session struct {
	ID            ID       `bson:"id" json:"Id"`
	UserName      string   `bson:"userName" json:"UserName"`
	UserID        ID       `bson:"uid" json:"UserId"`
	Admin         bool     `json:"-"`
	RemoteControl bool     `bson:"remoteControl" json:"SupportRemoteControl"`
	Commands      []string `bson:"commads" json:"SupportedCommands,array"`
	MediaTypes    []string `bson:"mediaType" json:"PlayableMediaTypes,array"`
	DeviceName    string   `bson:"deviceName" json:"DeviceName"`
	DeviceID      string   `bson:"deviceId" json:"DeviceId"`
	Client        string   `bson:"client" json:"Client"`
	ClientVersion string   `bson:"clientVersion" json:"ApplicationVersion"`
}

// SessionManager describes the interface that needs to be implemented
// by a session manager.
type SessionManager interface {
	// AddSession saves a new session.
	AddSession(session *Session)
	// GetSession retrieves a saved session.
	GetSession(id ID) *Session
	// RemoveSession removes a saved session from the manager.
	RemoveSession(uid, sessionID ID) bool
	// Starts up the manager.
	Run(saver SessionSaver) chan struct{}
}

// SessionSaver can save and get stored sessions from the database.
type SessionSaver interface {
	// GetSavedSessions retrieves saved sessions from database or disk.
	GetSavedSessions() ([]*Session, error)
	// SaveSessions saves the sessions to db or disk.
	SaveSessions([]*Session) error
}
