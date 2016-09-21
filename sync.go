package gomediacenter

// Sync is used to keep track of sync support for a file.
type Sync struct {
	SupportSync bool `json:"SupportsSync"`
	HasSyncJob  bool `json:"HashSyncJob"`
	IsSynced    bool `json:"IsSynced"`
}
