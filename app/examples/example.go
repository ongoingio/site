package examples

// Repository defines the interface to persist example entities.
type RepositoryInterface interface {
	Store(e *Example) error
	FindByAlias(e *Example) error
	FindAll(e *[]Example) error
	UpdateByAlias(e *Example) error
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
