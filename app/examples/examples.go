// Examples only needs the session.

package examples

//RepositoryInterface defines the interface
type RepositoryInterface interface {
	Store(example *Example) error
	FindByAlias(example *Example) error
	FindAll(examples *[]Example) error
	UpdateByAlias(example *Example) error
}

// Example represents an Example.
type Example struct {
	Name        string
	Alias       string
	Type        string
	Path        string
	Description string
	SHA         string
	Content     []Section
}

// TODO: Use Section from sidebyside.
type Section struct {
	Comment string
	Code    string
}
