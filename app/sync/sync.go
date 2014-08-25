// TODO: Ignore ignored files.
// TODO: Handle folders (and symlinks?) differently.

package sync

// TODO: Method on Examples or Example?
func generateAlias(name string) string {
	return url.QueryEscape(name)
}

// Prepare prepares the example with additional content.
func (example *Example) Prepare(content repository.Content) error {
	example.Type = content.Type
	example.Alias = generateAlias(content.Name)
	example.Name = content.Name
	example.SHA = content.SHA
	example.Path = content.Path

	// TODO: Extract description from file (via parse package).
	example.Description = "Some description..."

	// TODO: Prepare() in goroutine.
	err := content.Fetch()
	if err != nil {
		return err
	}
	example.Content = content.Content

	return nil
}

// Sync synchronizes the database with the Github repository.
func Sync() error {
	repo := repository.New("https://api.github.com/repos/ongoingio/examples/")
	repoContent, err := repo.Fetch()
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
