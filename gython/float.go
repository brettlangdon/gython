package gython

type Float struct {
	Value float64
}

func NewFloat(i float64) *Float {
	return &Float{
		Value: i,
	}
}

func (float *Float) object() {}
