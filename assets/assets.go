package assets

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"net/http"
	"net/url"
)

const (
	api      = "/api/v0/"
	ipfsHost = "localhost:5001"
	Max      = 46 // investigate byte limit
)

type Caches interface {
	Get(s string) []byte
}

type dataBlock []byte

type Cache struct {
	name   string
	origin string
	lock   sync.Mutex
	lru    *Lru
}

func NewCache() *Cache {
	c := &Cache{}
	c.lru = NewLru(Max)
	return c
}

type DummyCache struct {
	path  string
	local map[string][]byte
	lock  sync.Mutex
}

func (c *Cache) Req(path string, arg string) (resp *http.Response, err error) {
	u := url.URL{}
	u.Scheme = "http"
	u.Host = ipfsHost
	u.Path = api + path
	if arg != "" {
		val := url.Values{}
		val.Set("arg", arg)
		val.Set("encoding", "json")
		u.RawQuery = val.Encode()
	}
	fmt.Println(u.String())
	resp, err = http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

// like this /api/v0/name/resolve?arg=QmZXxbfUdRYi578pectWLFNFv5USQrsXdYAGeCsMJ6X8Zt&encoding=json
func (c *Cache) Resolve(name string) (data []byte, err error) {
	data, err = c.Get("name/resolve", name)
	return data, err
}

func (c *Cache) Ls(name string) (data []byte, err error) {
	data, err = c.Get("ls", name)
	return data, err
}

func (c *Cache) Diag() (data []byte, err error) {
	diag, err := c.Get("/diag/net", "")
	return diag, err
}

func (c *Cache) Cat(s string) (data dataBlock, err error) {
	val, ok := c.lru.Get(s)
	if !ok {
		// not in cache
		data, err = c.Get("cat", s)
		if err != nil {
			return nil, err
		}
		fmt.Println("add to cache", s)
		c.lru.Add(s, data)
		return data, nil
	}
	fmt.Println("in cache ", s)
	return val, nil
}

// get an object from the cache
func (c *Cache) Get(s string, a string) (data dataBlock, err error) {
	fmt.Println(s)
	resp, err := c.Req(s, a)
	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", data)
	return data, err
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
