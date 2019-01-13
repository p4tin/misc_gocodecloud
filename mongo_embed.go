package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	InsertedAt time.Time     `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time     `json:"last_update" bson:"last_update"`
}

type User struct {
	Entity    `bson:",inline"`
	Name      string    `json:"name" bson:"name"`
	BirthDate time.Time `json:"birth_date" bson:"birth_date"`
}

func main() {
	info := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Timeout:  60 * time.Second,
		Database: "test",
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//	c := session.DB("test").C("users")

	user := User{
		Entity:    Entity{"123456789098", time.Now().UTC(), time.Now().UTC()},
		Name:      "John Belushi",
		BirthDate: time.Date(1959, time.February, 28, 0, 0, 0, 0, time.UTC),
	}

	session.DB("test").C("users").Insert(user)
}
