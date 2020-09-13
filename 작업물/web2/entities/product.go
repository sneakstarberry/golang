package entities

import "fmt"

type Post struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (p Post) ToString() string {
	return fmt.Sprintf("id: %d\ntitle: %s\ncontent: %s\nimage: %s\n", p.Id, p.Title, p.Content, p.Image)
}
