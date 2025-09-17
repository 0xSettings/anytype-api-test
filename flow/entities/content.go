package entities

type Content struct {
	ID     string `json:"id,omitempty"`
	PageID string `json:"pageId"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body"`
}
