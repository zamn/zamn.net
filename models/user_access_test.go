package blog

import "testing"

func admin() *User {
	tempAdmin := &User{"zamn", "Adam", "Hamot", nil}
	tempAdmin.UserLevel = &Admin{tempAdmin}
	return tempAdmin
}

func TestAdminCreatePost(t *testing.T) {
	tempUser := admin()
	_, err := tempUser.UserLevel.CreatePost("Testing", "Hello World")
	if err != nil {
		t.Fail()
	}
}
