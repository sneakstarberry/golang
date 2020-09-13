package model

import (
	"database/sql"
	"net/http"

	"github.com/sneakstarberry/web2/entities"
)

type PostModel struct {
	Db *sql.DB
}

func (p PostModel) FindAll(r *http.Request) (post []entities.Post, err error) {
	rows, err := p.Db.Query("select * from post")
	if err != nil {
		return nil, err
	} else {
		var posts []entities.Post
		for rows.Next() {
			var id int64
			var title string
			var content string
			var image string
			err2 := rows.Scan(&id, &title, &content, &image)
			if err2 != nil {
				return nil, err2
			} else {
				hostURL := "http://" + r.Host + "/"
				post := entities.Post{
					Id:      id,
					Title:   title,
					Content: content,
					Image:   hostURL + image,
				}
				posts = append(posts, post)
			}
		}
		return posts, nil
	}
}

func (p PostModel) Create(post *entities.Post) error {
	result, err := p.Db.Exec("insert into post(title, content, image) values(?, ?, ?)", post.Title, post.Content, post.Image)
	if err != nil {
		return err
	} else {
		post.Id, _ = result.LastInsertId()
		return nil
	}
}
