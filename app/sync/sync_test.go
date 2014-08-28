package sync

import (
	"testing"

	"github.com/ongoingio/site/app/database"
	"github.com/ongoingio/site/app/examples"
	"github.com/ongoingio/site/app/repository"
)

type repositoryMock struct {
	URL string
}

func (repo *repositoryMock) Fetch(path string) (*repository.Content, []*repository.Content, error) {
	/*
	   content := &contentMock{
	       Type:     "file",
	       Encoding: "base64",
	       Name:     "README.md",
	       Path:     "README.md",
	       SHA:      "7962fa277d1c99417188f9fafe5ac3d575b22133",
	       URL:      "https://api.github.com/repos/ongoingio/examples/contents/README.md?ref=master",
	   }
	   contents := make([]contentMock, 1)
	   contents[0] = content
	   repos = &repository.Repository{URL: "foo/bar", Content: contents}
	*/

	return nil, nil, nil
}

func TestSyncNew(t *testing.T) {
	db := database.Connect("localhost/ongoing-test")
	defer db.Session.Close()
	err := db.DropDatabase()
	if err != nil {
		panic(err)
	}

	// Repository
	repo := &repositoryMock{
		URL: "/foo",
	}

	// Model
	examples.Register(db)

	// TODO: Refactor to use DI and not "global" M *Manager
	m := examples.M

	err = Sync(repo, m)
	if err != nil {
		t.Fatalf("Sync(): %v", err)
	}
}

func TestSyncChanged(t *testing.T) {
	// TODO
}

func TestSyncNone(t *testing.T) {
	// TODO
}
