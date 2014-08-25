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

func (repo *repositoryMock) Fetch(path string) {
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
}

func TestSync(t *testing.T) {
	session, err := database.Connect("localhost")
	if err != nil {
		panic(err)
	}

	repo := &repositoryMock{"foo/bar"}
	Register(session, repo)
	Sync()
}
