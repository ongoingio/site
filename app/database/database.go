package database

import "gopkg.in/mgo.v2"

// Session holds the database session.
var DB *mgo.Database

// Connect creates a database connection based on the given connection string
// and returns the database. The returned database struct holds the underlying
// session in its Session field.
// When the connection fails, the program panics. This is in favor of a explicit
// error return value with respect to the severity of the error.
func Connect(uri string) *mgo.Database {
	// See http://godoc.org/gopkg.in/mgo.v2#Dial
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	// With an empty name, mgo uses the db name from the connection string.
	// See http://godoc.org/gopkg.in/mgo.v2#Session.DB
	DB = session.DB("")

	return DB
}
