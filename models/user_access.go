package blog

import "time"

type Admin struct {
	User *User
}
type Editor struct {
	User *User
}
type StandardUser struct {
	User *User
}

type Access interface {
	CreatePost(title string, body string) (*Post, error)
	UpdatePost(*Post) error
	DeletePost(*Post) error
	ViewPost(*Post) error
	String() string
}

func (a Admin) CreatePost(title string, body string) (*Post, error) {
	// TODO: Implement mongodb
	return &Post{a.User, time.Now(), title, body, 0}, nil
}

func (a Admin) UpdatePost(*Post) error {
	return nil
}

func (a Admin) DeletePost(p *Post) error {
	return nil
}

func (a Admin) ViewPost(p *Post) error {
	return nil
}

func (a Admin) String() string {
	return "Admin"
}
