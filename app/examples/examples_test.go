package examples

import (
	"testing"
	"net/http/httptest"

	"github.com/ongoingio/site/app/repository"
)

func contentMock() repository.RepositoryContent {
	mock := repository.RepositoryContent{
		Type: "file",
		Encoding: "base64",
		Name: "README.md",
		Path: "README.md",
		Content: "R28gRXhhbXBsZXMKPT09PT09PT0KCkFuIGV4YW1wbGUgYSBkYXksIGtlZXBz\nIHRoZSBmcnVzdHJhdGlvbiBhd2F5Lgo=\n",
		SHA: "7962fa277d1c99417188f9fafe5ac3d575b22133",
		URL: "https://api.github.com/repos/ongoingio/examples/contents/README.md?ref=master",
	}

	return mock
}

func TestSync(t *testing.T) {

}