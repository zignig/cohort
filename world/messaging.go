package world

import (
	"encoding/json"
	"fmt"
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
