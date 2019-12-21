package main

import (
	"io/ioutil"

	"github.com/containous/yaegi/interp"
)

var plugins = make([]string, 0)

func loadPlugins() {
	dat, err := ioutil.ReadFile("plugins/plugin1.go")
	check(err)
	plugins = append(plugins, string(dat))

	dat, err = ioutil.ReadFile("plugins/plugin2.go")
	check(err)
	plugins = append(plugins, string(dat))
}

func main() {
	loadPlugins()
	i := interp.New(interp.Options{})
	for _, plugin := range plugins {
		i.Eval(plugin)
		v, _ := i.Eval("foo.Bar")
		bar := v.Interface().(func(string) string)
		r := bar("Kung")
		println(r)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
