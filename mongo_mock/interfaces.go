package mongo_mock

import (
	"gopkg.in/mgo.v2"
)

type DB interface {
	Dial(url string) (*mgo.Session, error)
	C(name string) *mgo.Collection
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
}

type DbSession interface {
	DB(name string) *mgo.Database
}