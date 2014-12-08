package assets

import (
	"fmt"
	"sync"
)

const api = "/api/v0/"

type Caches interface {
	Get(s string) []byte
}

type Cache struct {
	name   string
	origin string
	local  map[string][]byte
	lock   sync.Mutex
}

func NewCache() *Cache {
	c := &Cache{}
	return c
}

type DummyCache struct {
	path  string
	local map[string][]byte
	lock  sync.Mutex
}

// get an object from the cache
func (c *Cache) Get(s string) (data []byte) {
	fmt.Println(s)
	return data
}

type Item struct {
	Name string
	Hash string
	Size int64
}

type List struct {
	Hash  string
	Links []Item
}

type Listing struct {
	Objects []List
}
