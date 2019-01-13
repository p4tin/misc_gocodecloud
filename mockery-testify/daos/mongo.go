package daos

import (
	"fmt"

	"misc/mockery-testify"

	mgo "gopkg.in/mgo.v2"
)

// MongoDAL is an implementation of DataAccessLayer for MongoDB
type MongoDAL struct {
	session *mgo.Session
	dbName  string
}

// NewMongoDAL creates a MongoDAL
func NewMongoDAL(dbURI string, dbName string) (mockery_testify.DataAccessLayer, error) {
	session, err := mgo.Dial(dbURI)
	mongo := &MongoDAL{
		session: session,
		dbName:  dbName,
	}
	fmt.Println("Hello")
	return mongo, err
}

// c is a helper method to get a collection from the session
func (m *MongoDAL) c(collection string) *mgo.Collection {
	return m.session.DB(m.dbName).C(collection)
}

// Insert stores documents in mongo
func (m *MongoDAL) Insert(collectionName string, docs ...interface{}) error {
	return m.c(collectionName).Insert(docs)
}
