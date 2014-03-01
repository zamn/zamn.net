package blog

import "testing"

func admin() *User {
	return &User{"zamn", "Adam", "Hamot", &Admin{}}
}

func TestAdminCreatePost(t *testing.T) {
	tempUser := admin()
	_, err := tempUser.UserLevel.CreatePost("Testing", "Hello World")
	if err == nil {
		t.Fail()
	}
}
