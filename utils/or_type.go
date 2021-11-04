package utils

import (
	"encoding/json"
	"fmt"
)

type OR struct{}

type Demo1 struct {
	A int
}

type Demo2 struct {
	B int
}

type Demo3 struct{ C int }

type ORTypeDemo struct {
	OR
	demo1 *Demo1
	demo2 *Demo2
	demo3 *Demo3
}

func (o ORTypeDemo) MarshalJSON() ([]byte, error) {
	var notNil interface{}
	if o.demo1 != nil{
		notNil = o.demo1
	}
	if o.demo2 != nil{
		if notNil != nil{
			return nil, fmt.Errorf("not Support multiple not nil")
		}
		notNil = o.demo2
	}
	if o.demo3 != nil{
		if notNil != nil{
			return nil, fmt.Errorf("not Support multiple not nil")
		}
		notNil = o.demo3
	}
	return json.Marshal(notNil)
}

