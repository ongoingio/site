package sync

import (
	"testing"

	"github.com/ongoingio/site/app/database"
	"github.com/ongoingio/site/app/examples"
	"github.com/ongoingio/site/app/github"
)

type githubMock struct {
	URL string
}

func (repo *githubMock) Fetch(path string) (*github.Content, []*github.Content, error) {
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
	repo := &githubMock{
		URL: "/foo",
	}

	// Model
	manager := examples.New(db)

	err = Sync(repo, manager)
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
