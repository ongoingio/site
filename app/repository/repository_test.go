// TODO: Tests require consecutive order. Make them independent.

package repository

import (
    "testing"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "github.com/ongoingio/site/app/examples"
)

var c *mgo.Collection

func connect() {
    if c != nil {
        return
    }

    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }

    c = session.DB("ongoingio-test").C("examples")

    err = c.DropCollection()
    if err != nil {
        panic(err)
    }
}

func getExample() examples.Example {
    return examples.Example{
        Name: "Test",
        Alias: "test",
        Type: "file",
        Path: "test",
        Description: "Foobar",
        SHA: "123",
        Content: []examples.Section{{Comment: "a", Code: "aa"}},
    }
}

func TestStore(t *testing.T) {
    connect()
    e := getExample()

    repo := &Repository{Collection: c}
    repo.Store(&e)

    result := examples.Example{}
    err := c.Find(bson.M{"alias": "test"}).One(&result)
    if err != nil {
        t.Fatalf("mgo err: %v", err)
    }

    if result.Name != "Test" {
        t.Fatalf("name should be %s, is %s", "Test", result.Name)
    }
}

func TestFindByAlias(t *testing.T) {
    connect()
    e := getExample()

    err := c.Insert(&e)
    if err != nil {
        t.Fatalf("mgo err: %v", err)
    }

    repo := &Repository{Collection: c}
    result := &examples.Example{Alias: "test"}
    err = repo.FindByAlias(result)
    if result.Name != "Test" {
        t.Fatalf("name should be %s, is %s", "Test", result.Name)
    }
}

func TestInterface(t *testing.T) {
    repo := &Repository{}
    var i interface{} = repo
    if _, ok := i.(examples.RepositoryInterface); ok == false {
        t.Fatal("Repository does not fullfill the RepositoryInterface interface")
    }
}

func disconnect() {
    if c != nil {
        defer c.Database.Session.Close()
    }
}
