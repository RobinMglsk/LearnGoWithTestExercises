package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/robinmglsk/blogposts"
)

type StubFailingFS struct {}
func (s StubFailingFS) Open(name string) (fs.File, error){
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T){

	const(
		firstBody  = `Title: Post 1
Description: Description 1
Tags: tag1, tag2
---
Hello
World`
		secondBody  = `Title: Post 2
Description: Description 2
Tags: tag1, tag2
Some body
That I used
To know`
	)

	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs){
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "Description 1",  Tags: []string{"tag1", "tag2"}, Body: `Hello
World`})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %+v, want %+v", got, want)
	}
}