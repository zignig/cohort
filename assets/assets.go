package assets

import (
	"fmt"
	"sync"
)

type cache struct {
	name   string
	origin string
	local  map[string][]byte
	lock   sync.Mutex
}

// get an object from the cache
func (c *cache) Get(s string) (data []byte) {
	fmt.Println(s)
	return data
}
