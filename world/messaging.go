package world

import (
	"encoding/json"
	"fmt"
	"github.com/zignig/viewer/assets"
)

// message from player
type playMessage struct {
	Class   string          `json:"class"`
	Message json.RawMessage `json:"message"`
	Data    interface{}
}

// tagged "location"
type PosMessage struct {
	Pos  V3     `json:"pos"`
	Rot  E4     `json:"rot"`
	Uuid string `json:"uuid"`
}

// tagger "loader"
type LoaderMessage struct {
	Path string    `json:"path"`
	Pos  assets.V3 `json:"pos"`
	Rot  assets.E4 `json:"rot"`
}

// decodes play messages and returns objects into player loop

func (pm *playMessage) Decode(m []byte) {
	var dst interface{}
	err := json.Unmarshal(m, pm)
	if err != nil {
		fmt.Println("base decode ", err)
	}
	//fmt.Println("class ", pm.Class)
	switch pm.Class {
	case "location":
		{
			dst = &PosMessage{}
		}
	}
	err = json.Unmarshal(pm.Message, dst)
	pm.Data = dst
	if err != nil {
		fmt.Println("message error ", err)
	}
}

func Encode(i interface{}) (data []byte, err error) {
	pm := &playMessage{}
	switch i.(type) {
	default:
		pm.Class = "unknown"
	case LoaderMessage:
		pm.Class = "loader"
	case PosMessage:
		pm.Class = "location"
	}
	pm.Data, err = json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}
	data, err = json.Marshal(pm)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}
