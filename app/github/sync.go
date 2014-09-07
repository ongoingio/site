// Sync as a service coordinates the (one-way) synchronization of examples from Github to a local repository (data store).
//
// TODO: Ignore ignored files.
// TODO: Handle folders (and symlinks?) differently.

package github

import (
	"net/url"

	"gopkg.in/mgo.v2"

	"github.com/ongoingio/site/app/examples"
)

// generateAlias takes a string and returns a URL-safe representation of it.
func generateAlias(name string) string {
	return url.QueryEscape(name)
}

func prepare(repo Fetcher, content *Content) (*examples.Example, error) {
	example := &examples.Example{
		Type:  content.Type,
		Alias: generateAlias(content.Name),
		Name:  content.Name,
		SHA:   content.SHA,
		Path:  content.Path,
	}

	// TODO: Prepare() in goroutine.
	_, _, err := repo.Fetch(content.Path)
	if err != nil {
		return nil, err
	}

	// TODO: Parse content
	//example.Content = c.Content

	return example, nil
}

// Sync fetches examples from a remote repository (interface) and puts them into a local storage repository (interface).
func Sync(manager examples.RepositoryInterface, repo Fetcher) error {
	_, contents, err := repo.Fetch("/")
	if err != nil {
		return err
	}

	for _, c := range contents {
		result := &examples.Example{Path: c.Path}
		err = manager.FindByAlias(result)
		switch {
		case err == mgo.ErrNotFound:
			// Example not found, create new one.
			// log.Printf("DEBUG: Doc %s not found.", content.Path)
			// TODO: Alternative to passing the repo around?
			example, prerr := prepare(repo, c)
			if prerr != nil {
				return prerr
			}
			serr := manager.Store(example)
			if serr != nil {
				return serr
			}
		case err != nil:
			return err
		case c.SHA != result.SHA:
			// Example found, but content changed, update.
			// log.Printf("DEBUG: Doc %s found, but needs updating.", content.Path)
			example, prerr := prepare(repo, c)
			if prerr != nil {
				return prerr
			}
			serr := manager.UpdateByAlias(example)
			if serr != nil {
				return serr
			}
		default:
			// Example found, do nothing.
			// log.Printf("DEBUG: Doc %s found.", content.Path)
		}
	}

	return nil
}
