package blog

import (
	"time"

	"github.com/zamN/zamn.net/models"
)

func CreatePost(author *blog.User, title string, body string) (*blog.Post, error) {
	return &blog.Post{author, time.Now(), title, body, blog.EDITING}, nil
}
