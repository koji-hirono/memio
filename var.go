package memio

type Var struct {
	buf []byte
}

func NewVar(buf []byte) *Var {
	return &Var{buf: buf}
}

func (v *Var) Bytes() []byte {
	return v.buf
}

func (v *Var) Grow(n int) error {
	c := cap(v.buf)
	if n <= c {
		v.buf = v.buf[:n]
		return nil
	}
	if c == 0 {
		c = 64
	}
	for c < n {
		c <<= 1
	}
	buf := make([]byte, len(v.buf), c)
	copy(buf, v.buf)
	v.buf = buf[:n]
	return nil
}
