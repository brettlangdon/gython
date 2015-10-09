package gython

type _None struct{}

func (none *_None) object() {}

var None *_None = &_None{}
