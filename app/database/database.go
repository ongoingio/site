package database

import "gopkg.in/mgo.v2"

// Session holds the database session.
var Session *mgo.Session

// Connect creates a database connection based on the given URI and returns
// the session.
func Connect(uri string) (*mgo.Session, error) {
	session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	Session = session

	return session, nil
}
