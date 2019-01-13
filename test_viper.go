package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

// any approach to require this configuration into your program.
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

func main() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	data := viper.Get("clothing")
	t, ok := data.(map[string]interface{})
	if ok {
		fmt.Println(t["trousers"]) // this would be "steve"
	} else {
		fmt.Println("Not a map")
	}

}
