// Repository package is a service to get content from a Github repository.
//
// TODO: Rename /repository to /services?
// TODO: Could the repository just be another store/gateway? And sync an adapter between the two stores with the same interface?
//
// Usage:
//     repo := repository.New("http://url")
//     content, contents, err := repo.Fetch("root/or/path")
//

package github

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The Fetcher interface is implemented by repositories to allow fetching of files and file content.
type Fetcher interface {
	Fetch(path string) (*Content, []*Content, error)
}

// Repository represents a Github repository.
type Repository struct {
	URL string
}

// Content represents a content (file, dir, symlink) inside the repository.
type Content struct {
	Type     string `json:"type,omitempty"`
	Encoding string `json:"encoding,omitempty"`
	Name     string `json:"name,omitempty"`
	Path     string `json:"path,omitempty"`
	Content  string `json:"content,omitempty"`
	SHA      string `json:"sha,omitempty"`
	URL      string `json:"url,omitempty"`
}

func decode(content string) (string, error) {
	c, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}

	return string(c), nil
}

// Fetch fetches the content of a given path. This can either be the content of
// a single file when the path leads to a file, or the contents of a directory
// when the path leads to a directory. In order to make it possible to
// differentiate between the two, both contents are returned although one is
// always nil.
func (repo *Repository) Fetch(path string) (*Content, []*Content, error) {
	// TODO: Make sure to correctly (slash separators) construct a URL.
	resp, err := http.Get(repo.URL + path)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	// TODO: Does a streaming decoder perform better?

	// Try to decode a file content first.
	var fileContent *Content
	fileErr := json.Unmarshal(body, &fileContent)
	if fileErr == nil {
		content, err := decode(fileContent.Content)
		if err != nil {
			return nil, nil, err
		}
		fileContent.Content = content
		return fileContent, nil, fileErr
	}

	// If that failed, try to decode a directory content.
	var dirContent []*Content
	dirErr := json.Unmarshal(body, &dirContent)
	if err == nil {
		return nil, dirContent, dirErr
	}

	// And if both failed, assume another decoding error.
	return nil, nil, fmt.Errorf("decoding failed: %s / %s", fileErr, dirErr)
}

// New takes a repository URL and allocates and returns a new Repository.
func New(url string) *Repository {
	return &Repository{
		URL: url,
	}
}
