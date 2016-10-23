package gomediacenter

// Person is an actor, artist, director etc.
type Person struct {
	Name     string
	ID       ID
	Role     string
	Type     string
	ImageTag string
}

// PersonDTO is the struct encoded to JSON.
type PersonDTO struct {
	Name     string `json:"Name"`
	ID       string `json:"Id"`
	Role     string `json:"Role"`
	Type     string `json:"Type"`
	ImageTag string `json:"PrimaryImageTag"`
}

// PersonToDTO converts the Person struct to the DTO.
func PersonToDTO(p Person) PersonDTO {
	return PersonDTO{
		Name:     p.Name,
		ID:       p.ID.String(),
		Role:     p.Role,
		Type:     p.Type,
		ImageTag: p.ImageTag,
	}
}

// DTOToPerson converts the DTO struct to a person struct.
func DTOToPerson(dto PersonDTO) (Person, error) {
	id, err := IDFromString(dto.ID)
	if err != nil {
		return Person{}, err
	}
	return Person{
		Name:     dto.Name,
		ID:       id,
		Role:     dto.Role,
		Type:     dto.Type,
		ImageTag: dto.ImageTag,
	}, nil
}
