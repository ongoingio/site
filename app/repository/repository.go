package repository

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url string = "https://api.github.com/repos/ongoingio/examples/"

// RepositoryContent represents a file or directory inside the repository.
type RepositoryContent struct {
	Type     string `json:"type,omitempty"`
	Encoding string `json:"encoding,omitempty"`
	Name     string `json:"name,omitempty"`
	Path     string `json:"path,omitempty"`
	Content  string `json:"content,omitempty"`
	SHA      string `json:"sha,omitempty"`
}

var content []RepositoryContent

// Decode decodes the base64 encoded content of a file.
func (content *RepositoryContent) Decode() []byte {
	if content.Encoding != "base64" {
		panic("Cannot decode " + content.Name)
	}
	decoded, err := base64.StdEncoding.DecodeString(content.Content)
	if err != nil {
		panic(err)
	}
	return decoded
}

// Fetch the repository content from Github.
func Fetch() ([]byte, error) {
	res, err := http.Get(url + "contents/")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		// TODO: Real error type.
		return nil, fmt.Errorf("Unexpected status code %s", res.StatusCode)
	}

	return body, nil
}

// TODO: Rename to FetchRepostory() and combine with Fetch().
func Get() []RepositoryContent {
	return content
}

// FetchRepositoryContent takes an alias and fetches the content from the repository.
// TODO: Can Sha be used as an ID? Instead of an non-unique name or an alias that needs to be encoded/decoded every time.
func FetchRepositoryContent(name string) (RepositoryContent, error) {
	log.Printf("DEBUG: fetching content %s", name)
	// TODO: We currently use Name, but Github actually expects a Path.
	res, err := http.Get(url + "contents/" + name)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	// TODO: package globals (like content) always name collide with local variable. Is there a naming convention?
	file := RepositoryContent{}

	err = json.Unmarshal(body, &file)
	if err != nil {
		log.Fatal(err)
	}

	content, err := base64.StdEncoding.DecodeString(file.Content)
	file.Content = string(content)

	/*
		// TODO: Better way than using an empty struct? nil doesn't work...
		// TODO: Generate a meaningful error.
		return RepositoryContent{}, fmt.Errorf("Content not found")
	*/
	// TODO: Error handling?
	return file, nil
}

/*
	// TODO: RepositoryContent should be a map with the path as its key.

	// data holds the repository (content) locally inside this package.
	var data Repository

	// GetData provides a public function to access the package local repository data.
	func GetData() *data {
		return &data
	}
*/

// TODO: Create a new Repository type, attach all methods and initialize a new repository in init(), instead of having single functions around.
func init() {
	body, err := Fetch()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Dynamic length?
	content = make([]RepositoryContent, 10)

	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Fatal(err)
	}

	/*
		// Remove ignored files from list of examples.
		// TODO
		contentsCleaned := make([]Content, len(contents))

		for i, content := range contents {
			for _, ignore := range config.Ignore {
				if content.Name == ignore {
					continue
				}
			}
			contentsCleaned[i] = content
			contentsCleaned[i].Alias = url.QueryEscape(content.Name)
		}
	*/
}
