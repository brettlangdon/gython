package gython

type Tuple struct {
}

func (tuple *Tuple) object() {}

func NewTuple() *Tuple {
	return &Tuple{}
}
