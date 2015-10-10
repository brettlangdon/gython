package gython

type Bytes struct {
	value []byte
}

func NewBytes() *Bytes {
	return &Bytes{
		value: make([]byte, 0),
	}
}

func (bytes *Bytes) object() {}
func (bytes *Bytes) Append(b byte) {
	bytes.value = append(bytes.value, b)
}
func (bytes *Bytes) String() string {
	return string(bytes.value)
}
func (bytes *Bytes) Value() []byte {
	return bytes.value
}
