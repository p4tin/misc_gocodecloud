package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Test struct {
	A map[string]string
	B string
	C int
}

var fixture = Test{
	A: map[string]string{
		"test1": "data1",
	},
	B: "Test1",
	C: 3,
}

func main() {
	fmt.Printf("%+v\n", fixture)

	testFixture := Test{}
	copier.Copy(&testFixture, &fixture)
	delete(testFixture.A, "test1")

	fmt.Printf("%+v\n", fixture)
	fmt.Printf("%+v\n", testFixture)
	//
	//testFixture2 := Test{}
	//copier.Copy(&testFixture2, &fixture)
	//testFixture2.A = "Test3"
	//
	//fmt.Printf("%+v\n", fixture)
	//fmt.Printf("%+v\n", testFixture)
	//fmt.Printf("%+v\n", testFixture2)
}
