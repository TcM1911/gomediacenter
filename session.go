package gomediacenter

// Session holds all the data for an authenticated session.
type Session struct {
	ID            ID       `bson:"id" json:"Id"`
	UserName      string   `bson:"userName" json:"UserName"`
	UserID        string   `bson:"uid" json:"UserId"`
	Admin         bool     `json:"-"`
	RemoteControl bool     `bson:"remoteControl" json:"SupportRemoteControl"`
	Commands      []string `bson:"commads" json:"SupportedCommands,array"`
	MediaTypes    []string `bson:"mediaType" json:"PlayableMediaTypes,array"`
	DeviceName    string   `bson:"deviceName" json:"DeviceName"`
	DeviceID      string   `bson:"deviceId" json:"DeviceId"`
	Client        string   `bson:"client" json:"Client"`
	ClientVersion string   `bson:"clientVersion" json:"ApplicationVersion"`
}
