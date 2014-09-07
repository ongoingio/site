package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	file Content
	dir  []Content
)

func init() {
	file = Content{
		Name:     "README.md",
		Path:     "README.md",
		SHA:      "7962fa277d1c99417188f9fafe5ac3d575b22133",
		URL:      "https://api.github.com/repos/ongoingio/examples/contents/README.md",
		Type:     "file",
		Content:  "SnVzdCBhIHRlc3Qh",
		Encoding: "base64",
	}

	dir = []Content{
		{
			Name: "README.md",
			Path: "README.md",
			SHA:  "7962fa277d1c99417188f9fafe5ac3d575b22133",
			URL:  "https://api.github.com/repos/ongoingio/examples/contents/README.md",
			Type: "file",
		},
		{
			Name: ".gitignore",
			Path: ".gitignore",
			SHA:  "836562412fe8a44fa99a515eeff68d2bc1a86daa",
			URL:  "https://api.github.com/repos/ongoingio/examples/contents/.gitignore",
			Type: "file",
		},
	}
}

func createServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		var err error

		switch r.URL.String() {
		case "/":
			content, err = json.Marshal(dir)
			if err != nil {
				panic(err)
			}
		case "/file":
			content, err = json.Marshal(file)
			if err != nil {
				panic(err)
			}
		default:
			panic("URL not found.")
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(content))
	}))

	return ts
}

func TestFetchDir(t *testing.T) {
	ts := createServer()
	defer ts.Close()

	repo := New(ts.URL)
	content, contents, err := repo.Fetch("/")
	if err != nil {
		t.Fatalf("Fetch(): %v", err)
	}

	if content != nil {
		t.Fatalf("content should be nil, is %s", content)
	}

	if contents[0].Name != dir[0].Name {
		t.Fatalf("name of first file should be %s, is %s", dir[0].Name, contents[0].Name)
	}
}

func TestFetchFile(t *testing.T) {
	ts := createServer()
	defer ts.Close()

	repo := New(ts.URL)
	content, contents, err := repo.Fetch("/file")
	if err != nil {
		t.Fatalf("Fetch(): %v", err)
	}

	if contents != nil {
		t.Fatalf("contents should be nil, is %s", contents)
	}

	if content.Name != file.Name {
		t.Fatalf("name of file should be %s, is %s", file.Name, content.Name)
	}
}
