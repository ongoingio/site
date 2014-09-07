// TODO: Tests require consecutive order. Make them independent.

package examples

import (
	crand "crypto/rand"
	mrand "math/rand"
	"strconv"
	"strings"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var c *mgo.Collection

func connect() {
	if c == nil {
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		c = session.DB("ongoingio-test").C("examples")
	}

	err := c.DropCollection()
	if err != nil && err.Error() != "ns not found" {
		panic(err)
	}
}

func random(str_size int) string {
	alphanum := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	crand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func generateExample() Example {
	r := random(8)
	rl := strings.ToLower(r)
	rn := strconv.Itoa(mrand.Int())

	return Example{
		Name:        r,
		Alias:       rl,
		Type:        "file",
		Path:        rl,
		Description: "Foobar",
		SHA:         rn,
		Content:     []Section{{Comment: "a", Code: "aa"}},
	}
}

func TestStore(t *testing.T) {
	connect()
	e := generateExample()

	repo := &Repository{Collection: c}
	repo.Store(&e)

	result := Example{}
	err := c.Find(bson.M{"alias": e.Alias}).One(&result)
	if err != nil {
		t.Fatalf("mgo err: %v", err)
	}

	if result.Name != e.Name {
		t.Fatalf("name should be %s, is %s", e.Name, result.Name)
	}
}

func TestFindByAlias(t *testing.T) {
	connect()
	e := generateExample()

	err := c.Insert(&e)
	if err != nil {
		t.Fatalf("mgo err: %v", err)
	}

	repo := &Repository{Collection: c}
	result := &Example{Alias: e.Alias}
	err = repo.FindByAlias(result)
	if result.Name != e.Name {
		t.Fatalf("name should be %s, is %s", e.Name, result.Name)
	}
}

func TestInterface(t *testing.T) {
	repo := &Repository{}
	var i interface{} = repo
	if _, ok := i.(RepositoryInterface); ok == false {
		t.Fatal("Repository does not fullfill the RepositoryInterface interface")
	}
}

func TestFindAll(t *testing.T) {
	connect()
	e1 := generateExample()
	e2 := generateExample()

	err := c.Insert(&e1, &e2)
	if err != nil {
		t.Fatalf("mgo err: %v", err)
	}

	repo := &Repository{Collection: c}
	var results []Example
	err = repo.FindAll(&results)
	if results == nil {
		t.Fatal("results should not be nil")
	}
	if len(results) != 2 {
		t.Fatalf("results len should be 2, is %v", len(results))
	}
}

func TestUpdateByAlias(t *testing.T) {
	connect()
	e := generateExample()

	err := c.Insert(&e)
	if err != nil {
		t.Fatalf("mgo err: %v", err)
	}

	e.Name = "Foobar"

	repo := &Repository{Collection: c}

	err = repo.UpdateByAlias(&e)
	if err != nil {
		t.Fatalf("UpdateByAlias err: %v", err)
	}

	result := &Example{}
	err = c.Find(bson.M{"alias": e.Alias}).One(&result)
	if err != nil {
		t.Fatalf("mgo err: %v", err)
	}

	if result.Name != e.Name {
		t.Fatalf("Name should be %s, is %s", e.Name, result.Name)
	}
}

func disconnect() {
	if c != nil {
		defer c.Database.Session.Close()
	}
}
