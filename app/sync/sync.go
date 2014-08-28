// TODO: Ignore ignored files.
// TODO: Handle folders (and symlinks?) differently.

package sync

import (
	"net/url"

	"gopkg.in/mgo.v2"

	"github.com/ongoingio/site/app/examples"
	"github.com/ongoingio/site/app/repository"
)

// TODO: Method on Examples or Example?
func generateAlias(name string) string {
	return url.QueryEscape(name)
}

func prepare(repo repository.Fetcher, content *repository.Content) (*examples.Example, error) {
	example := &examples.Example{
		Type:  content.Type,
		Alias: generateAlias(content.Name),
		Name:  content.Name,
		SHA:   content.SHA,
		Path:  content.Path,
	}

	// TODO: Prepare() in goroutine.
	c, _, err := repo.Fetch(content.Path)
	if err != nil {
		return nil, err
	}
	example.Content = c.Content

	return example, nil
}

// Sync synchronizes the database with the Github repository.
func Sync(repo repository.Fetcher, manager examples.RepositoryInterface) error {
	_, contents, err := repo.Fetch("/")
	if err != nil {
		return err
	}

	for _, c := range contents {
		result := &examples.Example{Path: c.Path}
		err = manager.FindOne(result)
		switch {
		case err == mgo.ErrNotFound:
			// Example not found, create new one.
			// log.Printf("DEBUG: Doc %s not found.", content.Path)
			// TODO: Alternative to passing the repo around?
			example, prerr := prepare(repo, c)
			if prerr != nil {
				return prerr
			}
			serr := manager.Insert(example)
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
			serr := manager.Update(example)
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
