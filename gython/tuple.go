package gython

type Tuple struct {
	length *Float
	values []Object
}

func NewTuple(values ...Object) *Tuple {
	return &Tuple{
		length: NewFloat(float64(len(values))),
		values: values,
	}
}

func (tuple *Tuple) object()          {}
func (tuple *Tuple) Length() *Float   { return tuple.length }
func (tuple *Tuple) Values() []Object { return tuple.values }
