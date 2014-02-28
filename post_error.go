package blog

import (
	"fmt"

	"github.com/zamN/zamn.net/models"
)

type Action int

const (
	CREATE Action = 0
	DELETE Action = 1
	VIEW   Action = 2
)

type BlogError struct {
	Author blog.User
	Post   blog.Post
	Action Action
}

func (be *BlogError) Error() string {
	var action string

	switch be.Action {
	case CREATE:
		action = "CREATE"
	case DELETE:
		action = "DELETE"
	case VIEW:
		action = "VIEW"
	}

	return fmt.Sprintf("Error performing %s, User: %v, Post: %v", action, be.Author, be.Post)
}
