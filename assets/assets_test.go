package assets

import (
	"encoding/json"

	//"io/ioutil"
	//"net/http"
	"fmt"
	"strings"
	"testing"

	"github.com/zignig/cohort/util"
)

func TestTiles(t *testing.T) {
	conf := util.GetConfig("../universe.toml")
	tiles := conf.Tile
	c := NewCache()
	st, err := c.Ls(tiles)
	if err != nil {
		fmt.Println("FAIL resolve")
	}
	fmt.Print(string(st))
	tileList := &Listing{}
	json.Unmarshal(st, tileList)
	links := tileList.Objects[0].Links

	for i := range links {
		TileName := links[i].Name
		if strings.HasSuffix(TileName, "obj") {
			fmt.Println(TileName)
		}
	}
	fmt.Println(len(links), " tiles in hash")

}

func aTestCache(t *testing.T) {
	conf := util.GetConfig("../universe.toml")
	baseRef := conf.Ref
	fmt.Println(conf)
	c := NewCache()
	fmt.Println(c)
	//c.Diag()
	st, err := c.Resolve(baseRef)
	if err != nil {
		fmt.Println("FAIL resolve")
	}
	fmt.Println(st)
	data, err := c.Cat(st + "/" + conf.Path)
	if err != nil {
		fmt.Println("ls errror ", err)
	}
	fmt.Println(string(data))
	//fmt.Println(string(data))
	// TODO need to decode json stuff in assets ( dodj hack )
	//d, e := c.Ls(string(st[1 : len(st)-1]))
	//fmt.Println(string(d), e)

	// import export struct tests
	//export()

	d, err := c.Ls(st)
	if err != nil {
		fmt.Println("ls errror ", err)
	}
	fmt.Println(string(d))
	// broken hash
	d, err = c.Ls("Qmeq1j9dwd3xY4e6D6Qtrvvbr6DXF3diKLvbS2ApBb1T6j")
	if err != nil {
		fmt.Println("ls errror ", err)
	}
	fmt.Println(string(d))

	//c.Cat("QmTJK6iE6hhBXCYAReV9ftQVnY8eTkyWcQMF5cQiSyD2ty")

	//c.Ls("QmVyRrPEvAtTEDLKyEZWVMUwN9w3iJJxkN4uiCNNWoSyUQ")

	//c.Cat("QmTJK6iE6hhBXCYAReV9ftQVnY8eTkyWcQMF5cQiSyD2ty")
	export()
	//p := V3{}

	//dump(p)
}
