package main

import (
	"encoding/json"
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var session *r.Session

// Struct tags are used to map struct fields to fields in the database
type RtPerson struct {
	Id    string `gorethink:"id,omitempty"`
	Name  string `gorethink:"name"`
	Place string `gorethink:"place"`
}

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:32770",
		Database: "test",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// Create a table
	createTable()

	// Insert a record
	id := insertRecord()

	if id != "" {
		// Update a record
		updateRecord(id)
	}

	// Fetch one
	fetchOneRecord()

	// Record count
	recordCount()

	// Fetch all
	fetchAllRecords()

	if id != "" {
		// Delete a record
		deleteRecord(id)
	}
}

func createTable() {
	result, err := r.DB("test").TableCreate("people").RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}

	printStr("*** Create table result: ***")
	printObj(result)
	printStr("\n")
}

func insertRecord() string {
	var data = map[string]interface{}{
		"Name":  "David Davidson",
		"Place": "Somewhere",
	}

	result, err := r.Table("people").Insert(data).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	printStr("*** Insert result: ***")
	printObj(result)
	printStr("\n")

	return result.GeneratedKeys[0]
}

func updateRecord(id string) {
	var data = map[string]interface{}{
		"Name":  "Steve Stevenson",
		"Place": "Anywhere",
	}

	result, err := r.Table("people").Get(id).Update(data).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	printStr("*** Update result: ***")
	printObj(result)
	printStr("\n")
}

func fetchOneRecord() {
	cursor, err := r.Table("people").Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	var person interface{}
	cursor.One(&person)
	cursor.Close()

	printStr("*** Fetch one record: ***")
	printObj(person)
	printStr("\n")
}

func recordCount() {
	cursor, err := r.Table("people").Count().Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	var cnt int
	cursor.One(&cnt)
	cursor.Close()

	printStr("*** Count: ***")
	printObj(cnt)
	printStr("\n")
}

func fetchAllRecords() {
	rows, err := r.Table("people").Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read records into persons slice
	var persons []RtPerson
	err2 := rows.All(&persons)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	printStr("*** Fetch all rows: ***")
	for _, p := range persons {
		printObj(p)
	}
	printStr("\n")
}

func deleteRecord(id string) {
	result, err := r.Table("people").Get(id).Delete().Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	printStr("*** Delete result: ***")
	printObj(result)
	printStr("\n")
}

func printStr(v string) {
	fmt.Println(v)
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}
