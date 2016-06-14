package gomediacenter

import "gopkg.in/mgo.v2/bson"

type Session struct {
	Id            bson.ObjectId `bson:"_id" json:"Id"`
	UserName      string        `bson:"userName" json:"UserName"`
	UserId        string        `bson:"uid" json:"UserId"`
	RemoteControl bool          `bson:"remoteControl" json:"SupportRemoteControl"`
	Commands      []string      `bson:"commads" json:"SupportedCommands,array"`
	MediaTypes    []string      `bson:"mediaType" json:"PlayableMediaTypes,array"`
	DeviceName    string        `bson:"deviceName" json:"DeviceName"`
	DeviceId      string        `bson:"deviceId" json:"DeviceId"`
	Client        string        `bson:"client" json:"Client"`
	ClientVersion string        `bson:"clientVersion" json:"ApplicationVersion"`
}
