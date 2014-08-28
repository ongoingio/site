// Examples only needs the session.

package examples

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ManagerInterface defines the interface
type ManagerInterface interface {
	Insert(example *Example) error
	FindOne(example *Example) error
	FindAll(examples *[]Example) error
	Update(example *Example) error
	Remove(example *Example) error
}

// Manager TODO: Describe
type manager struct {
	Collection *mgo.Collection
}

// Example represents an Example.
type Example struct {
	Path        string
	Type        string
	Alias       string
	Name        string
	SHA         string
	Description string
	Content     string
}

// Insert creates a new example in the collection.
func (m *manager) Insert(example *Example) error {
	err := m.Collection.Insert(example)
	if err != nil {
		return err
	}

	return nil
}

// FindOne finds one document in the collection.
func (m *manager) FindOne(example *Example) error {
	// TODO: Why does a struct to Select work, but directly to Find (without Select) not?
	err := m.Collection.Find(nil).Select(example).One(example)
	if err != nil {
		return err
	}

	return nil
}

// FindAll finds all documents in the collection.
func (m *manager) FindAll(examples *[]Example) error {
	// TODO: Why does a struct to Select work, but directly to Find (without Select) not?
	err := m.Collection.Find(nil).All(examples)
	if err != nil {
		return err
	}

	return nil
}

// Update an existing example.
func (m *manager) Update(example *Example) error {
	query := bson.M{"path": example.Path}
	err := m.Collection.Update(query, example)
	if err != nil {
		return err
	}

	return nil
}

// Remove an example from the collection.
func (m *manager) Remove(example *Example) error {
	query := bson.M{"path": example.Path}
	err := m.Collection.Remove(query)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new manager.
func New(db *mgo.Database) *manager {
	return &manager{Collection: db.C("examples")}
}
