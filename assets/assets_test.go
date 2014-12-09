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
	c.Resolve("QmZXxbfUdRYi578pectWLFNFv5USQrsXdYAGeCsMJ6X8Zt")

}

//func TestTracker(t *testing.T) {
//	var url = string("http://thingtracker.net/example.tracker")
//	fmt.Println(url)
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	fmt.Println(string(body), err)
//	var tr ThingTracker
//	err2 := json.Unmarshal(body, &tr)
//	fmt.Println(err2, tr)
//	xml_data, err3 := json.MarshalIndent(tr, "", "\t")
//	fmt.Println(string(xml_data), err3)

//}
