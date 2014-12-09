package assets

import (
	//"encoding/json"

	//"io/ioutil"
	//"net/http"
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewCache()
	fmt.Println(c)
	//c.Diag()
	st, err := c.Resolve("QmZXxbfUdRYi578pectWLFNFv5USQrsXdYAGeCsMJ6X8Zt")
	if err != nil {
		fmt.Println("FAIL resolve")
	}
	fmt.Println(string(st))
	//c.Ls("QmUF6m34MsroeoR1atFHKGbFJLXYsT8othxf3LdsUXpGdt")
	c.Ls("QmZKzYD8cJanTipCniJbYu85iUC7xEaFQhpzWcquwJKaY7")

	//c.Ls("Qmeq1j9dwd3xYBe6D6Qtrvvbr6DXF3diKLvbS2ApBb1T6j")
	//c.Cat("QmTJK6iE6hhBXCYAReV9ftQVnY8eTkyWcQMF5cQiSyD2ty")

	//c.Ls("QmVyRrPEvAtTEDLKyEZWVMUwN9w3iJJxkN4uiCNNWoSyUQ")

	//c.Cat("QmTJK6iE6hhBXCYAReV9ftQVnY8eTkyWcQMF5cQiSyD2ty")
}
