package memio

type Fixed struct {
	buf []byte
}

func NewFixed(buf []byte) *Fixed {
	return &Fixed{buf: buf}
}

func (f *Fixed) Bytes() []byte {
	return f.buf
}

func (f *Fixed) Grow(n int) error {
	c := cap(f.buf)
	if n <= c {
		f.buf = f.buf[:n]
		return nil
	}
	return ErrNoMem
}
