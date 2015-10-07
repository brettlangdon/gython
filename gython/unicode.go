package gython

type Unicode struct {
}

func (unicode *Unicode) object() {}

func NewUnicode() *Unicode {
	return &Unicode{}
}
