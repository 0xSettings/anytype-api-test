package entities

type Page struct {
	ID      string `json:"id,omitempty"`
	SpaceID string `json:"spaceId"`
	Title   string `json:"title"`
}
