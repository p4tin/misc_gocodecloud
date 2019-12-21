package main

import (
	"fmt"
	"strings"
)

func catch() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		if strings.Contains(recoveryMessage.(string), "Serious") {
			panic(recoveryMessage)
		}
		fmt.Println(recoveryMessage)
	}
	fmt.Println("This is recovery function...")
}

func executePanic() {
	panic("This is Mild Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	defer catch()
	executePanic()
	panic("This is Serious Panic Situation")
	fmt.Println("Main block is executed completely...")
}
