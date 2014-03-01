package blog

import (
	"fmt"
)

type BlogError struct {
	Author User
	Post   Post
	Action string
}

func (be *BlogError) Error() string {
	return fmt.Sprintf("Error performing %s, User: %v, Post: %v", be.Action, be.Author, be.Post)
}
