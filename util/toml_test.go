package util

import (
	//"encoding/json"

	//"io/ioutil"
	//"net/http"

	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	conf := GetConfig("example.toml")
	fmt.Println("Config In")
	fmt.Printf("%v", conf)
	fmt.Println("Config Out")
	err := conf.SaveConfig()
	fmt.Println(err)
}
