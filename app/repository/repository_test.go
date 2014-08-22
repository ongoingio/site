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
		case "/contents":
			content, err = ioutil.ReadFile("./test/repository.json")
			if err != nil {
				panic(err)
			}
		case "/gobble":
			content, err = ioutil.ReadFile("./test/content.json")
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

func TestFetchRepository(t *testing.T) {
	ts := createServer()
	defer ts.Close()

	repo := New(ts.URL)
	err := repo.Fetch()
	if err != nil {
		t.Fatalf("Fetch(): %v", err)
	}

	if repo.Content[0].Name != ".gitignore" {
		t.Fatalf("name of first file should be %s, is %s", ".gitignore", repo.Content[0].Name)
	}
}

func TestFetchContent(t *testing.T) {
	ts := createServer()
	defer ts.Close()

	content := &Content{URL: ts.URL + "/gobble"}

	err := content.Fetch()
	if err != nil {
		t.Fatalf("Fetch(): %v", err)
	}

	if content.Content != "Just a test!" {
		t.Fatalf("content should be %s, is %s", "Just a test!", content.Content)
	}
}
