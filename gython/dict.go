package gython

import "fmt"

type Dict struct {
	values map[Object]Object
}

func NewDict() *Dict {
	return &Dict{
		values: make(map[Object]Object, 0),
	}
}

func (dict *Dict) object() {}

func (dict *Dict) Length() *Float {
	return NewFloat(float64(len(dict.values)))
}

func (dict *Dict) GetItem(key Object) (Object, error) {
	value, exists := dict.values[key]
	if exists == false {
		return nil, fmt.Errorf("Key error: %#v", key)
	}
	return value, nil
}

func (dict *Dict) SetItem(key Object, value Object) {
	dict.values[key] = value
}
