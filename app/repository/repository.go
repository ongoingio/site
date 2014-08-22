// TODO: Rename /repository to /services?
// Repository is a service package to communicate with a Github repository.
//
// Usage:
//
//     myRepo := repository.New("https://api.github.com/repos/ongoingio/examples/")
//     myRepo.Sync()

package repository

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
// ResponseFetcher defines an interface to fetch content from a given URL.
type ResponseFetcher interface {
	Fetch(url string) ([]byte, error)
}

// TODO: Document!
type Fetcher struct{}

// Fetch fetches the content from the given URL.
func (Fetcher) Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
*/

// Repository represents the Github repository.
type Repository struct {
	URL     string
	Content []Content
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

/*
func fetchAndParse(url string, value interface{}) {
	// TODO: Combine duplicated GET and JSON related code from FetchRepository() and FetchContent().
}
*/

// Fetch fetches the list of content from the repository.
// Example: https://api.github.com/repos/ongoingio/examples/contents
// See: https://developer.github.com/v3/repos/contents/
func (repo *Repository) Fetch() error {
	// TODO: Create function to construct an url from segments, like path.Join().
	resp, err := http.Get(repo.URL + "/contents")
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	// TODO: Does a streaming decoder make sense in this case?
	err = json.Unmarshal(body, &repo.Content)
	if err != nil {
		return err
	}

	return nil
}

// Fetch fetches a contents content.
func (content *Content) Fetch() error {
	res, err := http.Get(content.URL)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	// We only need the Content field...
	// TODO: Use anonymous struct with one field content instead.
	v := &struct{ Content string }{}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	c, err := base64.StdEncoding.DecodeString(v.Content)
	if err != nil {
		return err
	}
	content.Content = string(c)

	return nil
}

// New allocates and returns a new Repository.
func New(repo string) *Repository {
	return &Repository{
		URL: repo,
	}
}
