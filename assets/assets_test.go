package assets

import (
	//"encoding/json"

	//"io/ioutil"
	//"net/http"
	"fmt"
	"testing"

	"github.com/zignig/viewer/util"
)

func TestCache(t *testing.T) {
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
	fmt.Println(string(data))
	// TODO need to decode json stuff in assets ( dodj hack )
	//d, e := c.Ls(string(st[1 : len(st)-1]))
	//fmt.Println(string(d), e)

	// import export struct tests
	//export()

	c.Ls("Qmeq1j9dwd3xYBe6D6Qtrvvbr6DXF3diKLvbS2ApBb1T6j")
	//c.Cat("QmTJK6iE6hhBXCYAReV9ftQVnY8eTkyWcQMF5cQiSyD2ty")

	//c.Ls("QmVyRrPEvAtTEDLKyEZWVMUwN9w3iJJxkN4uiCNNWoSyUQ")

	//c.Cat("QmTJK6iE6hhBXCYAReV9ftQVnY8eTkyWcQMF5cQiSyD2ty")
}
