package sync

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2"

	"github.com/ongoingio/site/app/examples"
	"github.com/ongoingio/site/app/github"
)

type githubMock struct {
	URL string
}

func (repo *githubMock) Fetch(path string) (*github.Content, []*github.Content, error) {
	contentA := &github.Content{
		Type:     "file",
		Encoding: "base64",
		Name:     "File A.md",
		Path:     "File A.md",
		SHA:      "123",
	}
	contentB := &github.Content{
		Type:     "file",
		Encoding: "base64",
		Name:     "File B.md",
		Path:     "File B.md",
		SHA:      "456",
	}

	if path == "/" {
		contents := make([]*github.Content, 2)
		contents[0] = contentA
		contents[1] = contentB
		return nil, contents, nil
	}

	if path == "File A.md" {
		contentA.Content = "Content A"
		return contentA, nil, nil
	}

	if path == "File B.md" {
		contentB.Content = "Content B"
		return contentB, nil, nil
	}

	return nil, nil, fmt.Errorf("path %s not found", path)
}

type exampleStore struct {
	data []examples.Example
}

func (store *exampleStore) Insert(example *examples.Example) error {
	store.data = append(store.data, *example)
	return nil
}

func (store *exampleStore) FindOne(example *examples.Example) error {
	for _, ex := range store.data {
		if ex.Path == example.Path {
			example = &ex
			return nil
		}
	}
	return mgo.ErrNotFound
}

func (store *exampleStore) FindAll(examples *[]examples.Example) error {
	examples = &store.data
	return nil
}

func (store *exampleStore) Update(example *examples.Example) error {
	for i, ex := range store.data {
		if ex.Path == example.Path {
			store.data[i] = *example
		}
	}
	return nil
}

func (store *exampleStore) Remove(example *examples.Example) error {
	for i, ex := range store.data {
		if ex.Path == example.Path {
			// Delete index from slice.
			store.data = append(store.data[:i], store.data[i+1:]...)
		}
	}
	return nil
}

func TestSyncNew(t *testing.T) {
	github := &githubMock{
		URL: "/foo",
	}

	store := &exampleStore{}

	err := Sync(store, github)
	if err != nil {
		t.Fatalf("Sync(): %v", err)
	}

	if store.data[0].Name != "File A.md" {
		t.Fatalf("name is %s, should be %s", store.data[0].Name, "File A.md")
	}

	if store.data[0].Content != "Content A" {
		t.Fatalf("content is %s, should be %s", store.data[0].Content, "Content A")
	}
}

func TestSyncChanged(t *testing.T) {
	// TODO
}

func TestSyncNone(t *testing.T) {
	// TODO
}
