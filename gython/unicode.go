package gython

type Unicode struct {
	Value []byte
}

func (unicode *Unicode) object() {}

func NewUnicode(value []byte) *Unicode {
	return &Unicode{
		Value: value,
	}
}
