package blog

import "testing"

func TestCreatePost(t *testing.T) {
	result := CreatePost("hello world")
	if result != true {
		t.Fail()
	}
}
