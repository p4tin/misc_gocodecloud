package main

import (
	"errors"
	"fmt"
	"time"
)

type cacheData struct {
	value interface{}
	lastUpdateTime time.Time
	lastUpdateNode string
}

type cacheInfo struct {
	nodename string
	cache map[interface{}]cacheData
}

func NewCache(name string) *cacheInfo {
	cache := &cacheInfo{}
	cache.nodename = name
	cache.cache = make(map[interface{}]cacheData)
	return cache
}

func (c *cacheInfo)Set(key, value interface{}) {
	d := cacheData{
		value: value,
		lastUpdateTime: time.Now(),
		lastUpdateNode: c.nodename,
	}
	c.cache[key] = d
}

func (c *cacheInfo)Get(key interface{}) (interface{}, string, error) {
	if val, ok := c.cache[key]; !ok {
		return nil, c.nodename, errors.New("Key not in map")
	} else {
		return val.value, c.nodename, nil
	}
}
//
//func main() {
//	c := NewCache("Local")
//	c.Set("1", "2")
//	fmt.Println(c.Get("1"))
//	fmt.Println(c.Get("2"))
//	c.Set("1", "3")
//	c.Set("2", "4")
//	fmt.Println(c.Get("1"))
//	fmt.Println(c.Get("2"))
//}
