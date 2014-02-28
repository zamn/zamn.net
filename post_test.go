package blog

import (
	"testing"

	"github.com/zamN/zamn.net/models"
)

func Initialize() *blog.User {
	testUser := &blog.User{"zamn", "adam", "hamot", blog.ADMIN}
	return testUser
}

func TestCreatePost(t *testing.T) {
	user := Initialize()
	_, err := CreatePost(user, "Testing", "Hello World")
	if err != nil {
		t.Fail()
	}
}
