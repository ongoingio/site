package repository

import (
	"github.com/ongoingio/site/app/examples"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	Collection *mgo.Collection
}

func (r *Repository) Store(e *examples.Example) error {
	err := r.Collection.Insert(e)
	if err != nil {
	    return err
	}

	return nil
}

func (r *Repository) FindByAlias(e *examples.Example) error {
	err := r.Collection.Find(bson.M{"alias": e.Alias}).One(e)
    if err != nil {
    	return err
    }

	return nil
}

func (r *Repository) FindAll(e *[]examples.Example) error {
	return nil
}

func (r *Repository) UpdateByAlias(e *examples.Example) error {
	return nil
}
