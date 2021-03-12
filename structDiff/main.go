package main

import (
	"fmt"

	"gopkg.in/d4l3k/messagediff.v1"
)

type otherStruct struct {
	C  int
	Ar []int
}

type someStruct struct {
	A     int
	B     int
	Other []*otherStruct
}

type masterAStruct struct {
	someStruct []someStruct
}

func main() {
	x := otherStruct{C: 6}
	y := otherStruct{C: 6}
	z := otherStruct{C: 7}
	xx := otherStruct{C: 7}
	yy := otherStruct{C: 7}
	a := someStruct{1, 2, []*otherStruct{&x, &z}}
	b := someStruct{1, 2, []*otherStruct{&y, &xx}}
	m1 := masterAStruct{someStruct: []someStruct{a, b}}

	b.Other[1] = &yy
	m2 := masterAStruct{someStruct: []someStruct{b, a}}

	diff, equal := messagediff.PrettyDiff(m1, m2)
	fmt.Println(equal)
	fmt.Println(diff)

	/*
		equal = true
		diff = ""
	*/
}
