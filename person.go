package gomediacenter

// Person is the struct encoded to JSON.
type Person struct {
	Name     string `json:"Name"`
	ID       ID     `json:"Id"`
	Role     string `json:"Role"`
	Type     string `json:"Type"`
	ImageTag string `json:"PrimaryImageTag"`
	ImdbID   string `json:"Imdb"`
}
