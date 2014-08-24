// Examples depends on two services: the database session and a repository (interface).
// It probably is best to both inject them when the example model is created.

package examples

import (
	"testing"

	"github.com/ongoingio/site/app/repository"
)

/*
var contentMock1 = repository.Content{
	Type:     "file",
	Encoding: "base64",
	Name:     "README.md",
	Path:     "README.md",
	Content:  "R28gRXhhbXBsZXMKPT09PT09PT0KCkFuIGV4YW1wbGUgYSBkYXksIGtlZXBz\nIHRoZSBmcnVzdHJhdGlvbiBhd2F5Lgo=\n",
	SHA:      "7962fa277d1c99417188f9fafe5ac3d575b22133",
	URL:      "https://api.github.com/repos/ongoingio/examples/contents/README.md?ref=master",
}
*/

type repositoryMock struct {
	repository.Repository
}

type contentMock struct {
	repository.Content
}

func (repo *repositoryMock) Fetch() {
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

func (content *contentMock) Fetch() {
	/*
		content = &contentMock{
			Type:     "file",
			Encoding: "base64",
			Name:     "README.md",
			Path:     "README.md",
			Content:  "Content goes here...",
			SHA:      "7962fa277d1c99417188f9fafe5ac3d575b22133",
			URL:      "https://api.github.com/repos/ongoingio/examples/contents/README.md?ref=master",
		}
	*/
}

func TestSync(t *testing.T) {
	// TODO: Create database session.

	repository := &repositoryMock{repository.Repository{
		URL:     "foo/bar",
		Content: []contentMock{},
	}}

	Register(session, repository)
	Sync()
}
