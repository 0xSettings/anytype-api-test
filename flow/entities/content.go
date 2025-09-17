// entities/space.go
package entities

type Space struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// entities/page.go
package entities

type Page struct {
    ID      string `json:"id"`
    SpaceID string `json:"space_id"`
    Title   string `json:"title"`
}

// entities/content.go
package entities

type Content struct {
    ID     string `json:"id"`
    PageID string `json:"page_id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}
