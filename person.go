package gomediacenter

// Person is an actor, artist, director etc.
type Person struct {
	Name     string `json:"Name"`
	ID       ID     `json:"Id"`
	Role     string `json:"Role"`
	Type     string `json:"Type"`
	ImageTag string `json:"PrimaryImageTag"`
}
