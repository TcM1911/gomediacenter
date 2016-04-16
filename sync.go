package gomediacenter

type Sync struct {
	SupportSync bool `json:"SupportsSync"`
	HasSyncJob  bool `json:"HashSyncJob"`
	IsSynced    bool `json:"IsSynced"`
}
