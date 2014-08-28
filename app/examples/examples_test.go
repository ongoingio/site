// TODO: Tests require consecutive order. Make them independent.

package examples

import (
	"testing"

	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ongoingio/site/app/database"
)

func TestInsert(t *testing.T) {
	db := database.Connect("localhost/ongoing-test")
	defer db.Session.Close()

	// Drop database
	err := db.DropDatabase()
	if err != nil {
		panic(err)
	}

	manager := New(db)

	// TODO: M doesn't make sense.
	manager.Insert(&Example{
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

func TestFindOne(t *testing.T) {
	db := database.Connect("localhost/ongoing-test")
	defer db.Session.Close()

	manager := New(db)

	search := &Example{Name: "Foo.bar"}
	manager.FindOne(search)
	// TODO: Handle error.

	if search.Path != "Foo.bar" {
		t.Fatalf("path should be %s, is %s", "Foo.bar", search.Path)
	}
}

func TestFindAll(t *testing.T) {
	db := database.Connect("localhost/ongoing-test")
	defer db.Session.Close()

	manager := New(db)

	var results []Example
	manager.FindAll(&results)
	// TODO: Handle error

	if results[0].Name != "Foo.bar" {
		t.Fatalf("name of first should be %s, is %s", "Foo.bar", results[0].Name)
	}
}

func TestUpdate(t *testing.T) {
	db := database.Connect("localhost/ongoing-test")
	defer db.Session.Close()

	manager := New(db)

	doc := &Example{Name: "Baz", Path: "Foo.bar"}
	manager.Update(doc)
	// TODO: Handle error

	result := Example{}
	// TODO: Store collection name in model. Is currently hardcoded.
	err := db.C("examples").Find(bson.M{"path": "Foo.bar"}).One(&result)
	if err != nil {
		t.Fatalf("error in mgo: %v", err)
	}

	if result.Name != "Baz" {
		t.Fatalf("name should be %s, is %s", "Baz", result.Name)
	}
}

func TestRemove(t *testing.T) {
	db := database.Connect("localhost/ongoing-test")
	defer db.Session.Close()

	manager := New(db)

	err := manager.Remove(&Example{Path: "Foo.bar"})
	if err != nil {
		panic(err)
	}
}
