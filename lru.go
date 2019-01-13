package main

import (
	"github.com/golang/groupcache/lru"
	"github.com/hashicorp/memberlist"
	"fmt"
)
func main() {

	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.
	_, err = list.Join([]string{"10.27.10.224"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}

	cache := lru.New(5)
	fmt.Println(cache.MaxEntries)
	fmt.Println(cache.Len())
	cache.Add("1", 1)
	cache.Add("2", 2)
	cache.Add("3", 3)
	cache.Add("4", 4)
	cache.Add("5", 5)

	val, ok := cache.Get("1")
	fmt.Println(val, ok)
	val, ok = cache.Get("5")
	fmt.Println(val, ok)

	cache.Add("6", 6)

	fmt.Println(cache.Len())

	val, ok = cache.Get("1")
	fmt.Println(val, ok)
	val, ok = cache.Get("5")
	fmt.Println(val, ok)
	val, ok = cache.Get("6")
	fmt.Println(val, ok)
}
