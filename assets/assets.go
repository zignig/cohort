package assets

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"net/http"
	"net/url"
)

const (
	api      = "/api/v0/"
	ipfsHost = "localhost:5001"
	Max      = 600 // investigate byte limit
)

type Caches interface {
	Get(s string) []byte
}

type dataBlock []byte

//type Ref map[string]string
type Ref struct {
	Key     string
	Message string
}

type Cache struct {
	name      string
	origin    string
	lock      sync.Mutex
	nameCache map[string]string
	nameLock  sync.Mutex
	lru       *Lru
}

func NewCache() *Cache {
	c := &Cache{}
	c.nameCache = make(map[string]string)
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
	//TODO need to parse and return http status
	fmt.Println(u.String())
	resp, err = http.Get(u.String())
	if resp.StatusCode != 200 {
		return resp, errors.New(resp.Status)
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

// like this /api/v0/name/resolve?arg=QmZXxbfUdRYi578pectWLFNFv5USQrsXdYAGeCsMJ6X8Zt&encoding=json
func (c *Cache) Resolve(name string) (ref string, err error) {
	val, ok := c.nameCache[name]
	if ok {
		fmt.Println("in name cache")
		return val, err
	}
	data, err := c.Get("name/resolve", name)
	if err != nil {
		fmt.Println("resolve error ", err)
		return "", err
	}
	refObj := &Ref{}
	fmt.Println("start unmarshall")
	merr := json.Unmarshal(data, &refObj)
	fmt.Println(refObj)
	if merr != nil {
		fmt.Println("unmarshall error ", merr)
		return "", err
	}
	if refObj.Key == "" {
		fmt.Println("key error ", merr)
		return "", err
	}
	ref = refObj.Key
	c.nameLock.Lock()
	fmt.Println("add name ", name, " to cache")
	c.nameCache[name] = ref
	c.nameLock.Unlock()
	return ref, err
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
			fmt.Println("cat error ", err, data)
			return data, err
		}
		fmt.Println("add to cache", s)
		c.lru.Add(s, data)
		return data, nil
	}
	//fmt.Println("in cache ", s)
	return val, nil
}

// get an object from the cache
func (c *Cache) Get(s string, a string) (data dataBlock, err error) {
	fmt.Println(s)
	resp, err := c.Req(s, a)
	if err != nil {
		fmt.Println(err)
		return dataBlock{}, err
	}
	fmt.Println(resp.Status)
	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return dataBlock{}, err
	}
	//fmt.Printf("%s", data)
	return data, err
}

func (c *Cache) Listing(path string) (items []string, err error) {
	resp, err := c.Ls(path)
	li := &Listing{}
	if err != nil {
		return items, err
	}
	fmt.Println("start unmarshall")
	merr := json.Unmarshal(resp, &li)
	if merr != nil {
		fmt.Println("Unmarshall error ", err)
		return items, merr
	}
	for _, it := range li.Objects[0].Links {
		//fmt.Println("listing ", i, it)
		items = append(items, it.Name)
	}
	return items, nil
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
