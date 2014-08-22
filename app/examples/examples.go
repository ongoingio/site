// TODO: Ignore ignored files.
// TODO: Handle folders differently.

package examples

import (
	_"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/url"

	"github.com/ongoingio/site/app/database"
	"github.com/ongoingio/site/app/repository"
)

var collection *mgo.Collection

type Example struct {
	Path        string
	Type        string
	Alias       string
	Name        string
	SHA         string
	Description string
	Content     string
}

/*
// TODO: Possible to List() as a method?
func List() []Example {
	return data
}
*/

/*
// TODO: Return a pointer or an error? Most ORM seem to be using `err := Find(&example, "name-1")`
func FindByName(name string) (Example, error) {
	for _, e := range data {
		if e.Name == name {
			return e, nil
		}
	}

	return Example{}, fmt.Errorf("%s not found", name)
}
*/

// TODO: Method on Examples or Example?
func generateAlias(name string) string {
	return url.QueryEscape(name)
}

func (example *Example) Save() error {
	err := collection.Insert(example)
	if err != nil {
		return err
	}

	return nil
}

func (example *Example) UpdateByPath(path string) error {
	colQuerier := bson.M{"path": path}
	change := bson.M{"$set": example}
	err := collection.Update(colQuerier, change)
	if err != nil {
		return err
	}

	return nil
}

func (example *Example) Prepare(content repository.RepositoryContent) error {
	example.Type = content.Type
	example.Alias = generateAlias(content.Name)
	example.Name = content.Name
	example.SHA = content.SHA
	example.Path = content.Path

	// TODO: Extract description from file (via parse package).
	example.Description = "Some description..."

	// TODO: Prepare() in goroutine.
	err := content.FetchContent()
	if err != nil {
		return err
	}
	example.Content = content.Content

	return nil
}

func Sync() error {
	repo := repository.New("https://api.github.com/repos/ongoingio/examples/")
	repoContent, err := repo.FetchRepository()
	if err != nil {
		log.Fatal(err)
	}

	for _, content := range repoContent.Content {
		result := Example{}
		err = collection.Find(bson.M{"path": content.Path}).One(&result)
		switch {
		case err == mgo.ErrNotFound:
			log.Printf("DEBUG: Doc %s not found.", content.Path)
			example := &Example{}
			example.Prepare(content)
			example.Save()
		case err != nil:
			return err
		case content.SHA != result.SHA:
			log.Printf("DEBUG: Doc %s found, but needs updating.", content.Path)
			example := &Example{}
			example.Prepare(content)
			example.UpdateByPath(content.Path)
		default:
			log.Printf("DEBUG: Doc %s found.", content.Path)
		}
	}

	return nil
}

func Register() {
	collection = database.Session.DB("ongoingio").C("examples")
}
