package main

import (
	"fmt"
	"reflect"

	"github.com/aerogo/nano"
)

type Cart struct {
}

var colName = reflect.TypeOf(Cart{})

var types = []interface{}{
	colName,
}

func main() {
	node := nano.New(nano.Configuration{})
	defer node.Close()
	defer node.Clear()

	node.Namespace("ecomm-api").RegisterTypes(types...)
	fmt.Printf("%v - %s - %+v\n", node.IsServer(), colName, node.Namespace("ecomm-api").Collection(colName.String()))
}
