package gython

type Bytes struct {
}

func (bytes *Bytes) object() {}

func NewBytes() *Bytes {
	return &Bytes{}
}
