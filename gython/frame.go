package gython

type Frame struct {
}

func NewFrame() *Frame {
	return &Frame{}
}

func (frame *Frame) object() {}
