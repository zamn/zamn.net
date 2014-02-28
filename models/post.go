package blog

import "time"

type View int

const (
	PUBLIC  View = 0
	EDITING View = 1
	HIDDEN  View = 2
)

type Post struct {
	Author    *User
	DateTime  time.Time
	Title     string
	Body      string
	ViewLevel View
}
