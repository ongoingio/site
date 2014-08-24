package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		var err error

		switch r.URL.String() {
		case "/":
			content, err = ioutil.ReadFile("./test/dir.json")
			if err != nil {
				panic(err)
			}
		case "/file":
			content, err = ioutil.ReadFile("./test/file.json")
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

	if contents[0].Name != ".gitignore" {
		t.Fatalf("name of first file should be %s, is %s", ".gitignore", contents[0].Name)
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

	if content.Name != "README.md" {
		t.Fatalf("name of file should be %s, is %s", "README.md", content.Name)
	}
}
