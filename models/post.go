package blog

import "time"

type Post struct {
	Author   *User
	DateTime time.Time
	Title    string
	Body     string
	Hidden   int
}
