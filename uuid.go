package main

import (
	"fmt"

	"github.com/twinj/uuid"
)

func main() {
	fmt.Println(uuid.NewV4().String())
}
