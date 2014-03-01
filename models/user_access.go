package blog

type Admin struct{}
type Editor struct{}
type StandardUser struct{}

type Access interface {
	CreatePost(title string, body string) (*Post, error)
	DeletePost(*Post) error
	ViewPost(*Post) error
	String() string
}

func (a Admin) CreatePost(title string, body string) (*Post, error) {
	return nil, nil
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
