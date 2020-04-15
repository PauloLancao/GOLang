package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	cache.mapping["cache1"] = "cache1"
	cache.mapping["cache1"] = "cache2"

	fmt.Println(lookup("cache1"))
	fmt.Println(lookup("cache2"))
}
