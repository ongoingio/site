// Examples only needs the session.

package examples

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Manager TODO: Describe
type Manager struct {
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
func (m *Manager) Insert(example *Example) error {
	err := m.Collection.Insert(example)
	if err != nil {
		return err
	}

	return nil
}

// FindOne finds one document in the collection.
func (m *Manager) FindOne(example *Example) error {
	// TODO: Why does a struct to Select work, but directly to Find (without Select) not?
	err := m.Collection.Find(nil).Select(example).One(example)
	if err != nil {
		return err
	}

	return nil
}

// FindAll finds all documents in the collection.
func (m *Manager) FindAll(examples *[]Example) error {
	// TODO: Why does a struct to Select work, but directly to Find (without Select) not?
	err := m.Collection.Find(nil).All(examples)
	if err != nil {
		return err
	}

	return nil
}

// Update an existing example.
func (m *Manager) Update(example *Example) error {
	query := bson.M{"path": example.Path}
	err := m.Collection.Update(query, example)
	if err != nil {
		return err
	}

	return nil
}

// Remove an example from the collection.
func (m *Manager) Remove(example *Example) error {
	query := bson.M{"path": example.Path}
	err := m.Collection.Remove(query)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new manager.
func New(db *mgo.Database) *Manager {
	return &Manager{Collection: db.C("examples")}
}
