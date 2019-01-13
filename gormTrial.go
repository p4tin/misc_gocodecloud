package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type Status struct {
	ID       	uint
	Text 		string
	UserID  	uint
	StatusID 	uint		//Parent Status (top level == nil parent)
}

//type Address struct {
//	Line1 		string
//	Line2 		string
//	Line3 		string
//	City 		string
//	State 		string
//	Country 	string
//	PostalCode 	string
//}

type Profile struct {
	FirstName 	string
	LastName 	string
	RelationShip 	string
	Occupation 	string
//	Address 	Address
}

type GUser struct {
	gorm.Model
	Email 		string 		`gorm:"unique"`
	Password 	string
	Profile 	Profile
	Statuses    	[]Status
	Requests 	[]GUser
	Friends 	[]GUser
}

var db *gorm.DB

func (u *GUser)Create() {
	db.Create(u)
}


type BlankLogger struct {}
func (l *BlankLogger)Print(v ...interface{}) {}

func main() {
	var err error
	db, err = gorm.Open("sqlite3", "db/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.SetLogger(&BlankLogger{})

	// Migrate the schema
	db.AutoMigrate(&Status{}, &Profile{}, &GUser{})

	// Create
	prof1 := Profile{FirstName: "Paul", LastName: "Fortin"}
	user1 := &GUser{Email: "paul@cp-soft.com", Password: "password", Profile: prof1, }
	if dbc := db.Create(user1); dbc.Error != nil {
		fmt.Println("--> " + dbc.Error.Error())
	}
	prof2 := Profile{FirstName: "Joey", LastName: "Hayward"}
	user2 := &GUser{Email: "joe@yahoo.com", Password: "password", Profile: prof2, }
	if dbc := db.Create(user2); dbc.Error != nil {
		fmt.Println("--> " + dbc.Error.Error())
	}

	// Read
	var user GUser
	db.First(&user, 1) // find product with id 1
	db.First(&user, "email = ?", "paul@cp-soft.com") // find profile with email paul...

	// Update - update product's price to 2000
	db.Model(&user).Update("Password", "test12345")

	// Delete - delete product
	//db.Delete(&profile)


	db.First(&user, 2).Related(&user.Profile)

	fmt.Printf("%#v\n", user)
}