package world

import (
	"encoding/json"
	"fmt"
	"github.com/zignig/cohort/assets"
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

// tagger "floor"
type FloorMessage struct {
	Pos  assets.V3 `json:"Pos"`
	Size int       `json:"Size"`
}

// tagger "tile"
type TileMessage struct {
	Pos    assets.V3 `json:"Pos"`
	Name   string
	Ref    string
	Rotate int
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
	case *FloorMessage:
		fmt.Println(v)
		ps.Class = "floor"
	case *LoaderMessage:
		fmt.Println(v)
		ps.Class = "loader"
	case *PosMessage:
		fmt.Println(v)
		ps.Class = "location"
	case *TileMessage:
		fmt.Println(v)
		ps.Class = "tile"
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
