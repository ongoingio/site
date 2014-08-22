package database

import (
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var Session *mgo.Session

func Connect() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	Session = session
}
