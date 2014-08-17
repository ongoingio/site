package model

import (
	"log"
	"net/url"

	"github.com/ongoingio/site/app/repository"
)

type Example struct {
	Type        string
	Alias       string
	Name        string
	Description string
	Content     string
}

type Examples []Example

func List() []Example {
	return examples
}

// FindByName gets the content of a single file from the repository.
// TODO: Return a pointer or an error? Most ORM seem to be using `err := Find(&example, "name-1")`
// TODO: GetByAlias vs FindByAlias?
// TODO: Should probably be FindByPath().
func FindByName(name string) (Example, bool) {
	content, err := repository.FetchRepositoryContent(name)
	if err != nil {
		// TODO: Don't use bool...
		log.Fatal(err)
		return Example{}, false
	}

	// Transform from RepositoryContent to Example.
	example := Example{
		Type:        content.Type,
		Alias:       generateAlias(content.Name),
		Name:        content.Name,
		Description: "Some description goes here...",
		Content:     content.Content,
	}

	return example, true
}

// TODO: Method on Examples or Example?
func generateAlias(name string) string {
	return url.QueryEscape(name)
}

// TODO: Alternative way to hold data inside a package? No one else (ORMs etc.) seems to be doing it that way.
var examples Examples

// TODO: Since we already need the decription on the list page, we need the full repository with all contents fetched.
func init() {
	/*
		// Dummy content.
		examples = append(examples,
			Example{"file", "name-1", "Name 1", "Description 1", "Content 1"},
			Example{"file", "name-2", "Name 2", "Description 2", "Content 2"},
			Example{"dir", "name-3", "Name 3", "Description 3", "Content 3"})
	*/
	content := repository.Get()

	// TODO: Cleaner way to transform RepositoryContent into Example?
	for _, c := range content {
		example := Example{Type: c.Type, Name: c.Name, Description: "Some Description"}
		example.Alias = generateAlias(c.Name)
		examples = append(examples, example)
	}
}
