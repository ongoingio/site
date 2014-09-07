package examples

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	Collection *mgo.Collection
}

func (r *Repository) Store(e *Example) error {
	err := r.Collection.Insert(e)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindByAlias(e *Example) error {
	err := r.Collection.Find(bson.M{"alias": e.Alias}).One(e)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindAll(e *[]Example) error {
	err := r.Collection.Find(bson.M{}).All(e)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateByAlias(e *Example) error {
	err := r.Collection.Update(bson.M{"alias": e.Alias}, e)
	if err != nil {
		return err
	}

	return nil
}
