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
	data    interface{}
}

// message to player
type playSend struct {
	Class   string      `json:"class"`
	Message interface{} `json:"message"`
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
	Pos  assets.V3 `json:"Pos"`
	Rot  assets.E4 `json:"Rot"`
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
	pm.data = dst
	if err != nil {
		fmt.Println("message error ", err)
	}
}

func Encode(i interface{}) (data []byte, err error) {
	ps := &playSend{}
	switch v := i.(type) {
	case *LoaderMessage:
		fmt.Println(v)
		ps.Class = "loader"
	case *PosMessage:
		fmt.Println(v)
		ps.Class = "location"
	default:
		fmt.Println(v)
		ps.Class = "infiniteawesome"
	}

	ps.Message = i
	data, err = json.MarshalIndent(ps, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}
