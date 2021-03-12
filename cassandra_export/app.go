package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

type Person struct {
	id        string
	firstName string
	lastName  string
}

func main() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "example"

	session, err := cluster.CreateSession()
	checkerr(err)
	defer session.Close()

	fmt.Println("Getting all People")
	//var peeps []Person
	//m := map[string]interface{}{}
	query := "SELECT id, firstname, lastname FROM example.person"

	statement := session.Query(query)
	iter := statement.Iter()
	var id, fname, lname string
	for iter.Scan(&id, &fname, &lname) {
		fmt.Printf("%s - %s - %s\n", id, fname, lname)
		//peeps = append(peeps, Person{
		//	id:        m["id"].(string),
		//	firstName: m["first_name"].(string),
		//	lastName:  m["last_name"].(string),
		//})
		//m = map[string]interface{}{}
	}
	err = iter.Close()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", peeps)
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}

}
