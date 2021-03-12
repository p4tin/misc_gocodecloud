package main

import (
	"fmt"
	"math/rand"
	"strconv"
	ti "time"

	"github.com/wawandco/fako"
)

type FakeUser struct {
	LastName  string `fako:"last_name"`
	FirstName string `fako:"first_name"`
	Name      string `fako:"full_name"`
	Username  string `fako:"user_name"`
	Domain    string `fako:"domain_name"`
	Email     string `fako:"email_address"` //Notice the fako:"email_address" tag
	Phone     string `fako:"phone"`
	Password  string `fako:"simple_password"`
	Address   string `fako:"street_address"`
}

type Sku struct {
	Brand       string `fako:"brand"`
	ProductName string `fako:"product"`
	Size        string `fako:"size"`
	Color       string `fako:"color"`
}

func main() {
	fako.Register("size", func() string {
		s1 := rand.NewSource(ti.Now().UnixNano())
		r1 := rand.New(s1)
		return strconv.Itoa(r1.Intn(10))
	})

	var user FakeUser
	fako.Fill(&user)

	fmt.Printf("%#v\n", user)

	var sku Sku
	fako.Fill(&sku)

	fmt.Printf("%#v\n", sku)

	var userWithOnlyEmail FakeUser
	fako.FillOnly(&userWithOnlyEmail, "Email")
	//This will fill all only the email

	var userWithoutEmail FakeUser
	fako.FillExcept(&userWithoutEmail, "Email")
	//This will fill all the fields except the email

}
