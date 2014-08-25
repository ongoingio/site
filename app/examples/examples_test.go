package examples

import (
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestInsert(t *testing.T) {
	// TODO: Create a new session in every test?
	// TODO: Refactor out into function
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db := session.DB("ongoing-test")
	Register(db)

	// Drop database
	err = session.DB("ongoing-test").DropDatabase()
	if err != nil {
		panic(err)
	}

	// TODO: M doesn't make sense.
	M.Insert(&Example{
		Path:    "Foo.bar",
		Name:    "Foo.bar",
		Content: "Foo bar baz [...]",
	})

	result := Example{}
	// TODO: Store collection name in model. Is currently hardcoded.
	err = db.C("examples").Find(bson.M{"path": "Foo.bar"}).One(&result)
	if err != nil {
		t.Fatalf("error in mgo: %v", err)
	}

	if result.Name != "Foo.bar" {
		t.Fatalf("name should be %s, is %s", "Foo.bar", result.Name)
	}
}

// TODO: Tests require the insert first. Make them all independent?
func TestFindOne(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db := session.DB("ongoing-test")
	Register(db)

	search := &Example{Name: "Foo.bar"}
	M.FindOne(search)
	// TODO: Handle error.

	if search.Path != "Foo.bar" {
		t.Fatalf("path should be %s, is %s", "Foo.bar", search.Path)
	}
}

func TestFindAll(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db := session.DB("ongoing-test")
	Register(db)

	var results []Example
	M.FindAll(&results)
	// TODO: Handle error

	if results[0].Name != "Foo.bar" {
		t.Fatalf("name of first should be %s, is %s", "Foo.bar", results[0].Name)
	}
}

func TestUpdate(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db := session.DB("ongoing-test")
	Register(db)

	doc := &Example{Name: "Baz", Path: "Foo.bar"}
	M.Update(doc)
	// TODO: Handle error

	result := Example{}
	// TODO: Store collection name in model. Is currently hardcoded.
	err = db.C("examples").Find(bson.M{"path": "Foo.bar"}).One(&result)
	if err != nil {
		t.Fatalf("error in mgo: %v", err)
	}

	if result.Name != "Baz" {
		t.Fatalf("name should be %s, is %s", "Baz", result.Name)
	}
}

func TestRemove(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db := session.DB("ongoing-test")
	Register(db)

	err = M.Remove(&Example{Path: "Foo.bar"})
	if err != nil {
		panic(err)
	}
}
