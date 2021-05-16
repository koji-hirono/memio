package memio

import (
	"testing"
)

func TestFixed_Grow(t *testing.T) {
	t.Run("2 -> 2byte", func(t *testing.T) {
		p := make([]byte, 2)
		m := NewFixed(p)
		err := m.Grow(2)
		if err != nil {
			t.Fatal(err)
		}
		buf := m.Bytes()
		buf[0] = byte(0xcd)
		buf[1] = byte(0xab)

		buf = m.Bytes()
		if buf[0] != byte(0xcd) {
			t.Errorf("want %x; but got %x\n", byte(0xcd), buf[0])
		}
		if buf[1] != byte(0xab) {
			t.Errorf("want %x; but got %x\n", byte(0xab), buf[1])
		}
		if cap(buf) != 2 {
			t.Errorf("want %v; but got %v\n", 2, cap(buf))
		}
		if len(buf) != 2 {
			t.Errorf("want %v; but got %v\n", 2, len(buf))
		}
	})
	t.Run("out of memory", func(t *testing.T) {
		p := make([]byte, 2)
		m := NewFixed(p)
		err := m.Grow(3)
		if err != ErrNoMem {
			t.Errorf("want %v; but got %v\n", ErrNoMem, err)
		}
		buf := m.Bytes()
		if buf[0] != byte(0x0) {
			t.Errorf("want %x; but got %x\n", byte(0x0), buf[0])
		}
		if buf[1] != byte(0x0) {
			t.Errorf("want %x; but got %x\n", byte(0x0), buf[1])
		}
		if cap(buf) != 2 {
			t.Errorf("want %v; but got %v\n", 2, cap(buf))
		}
		if len(buf) != 2 {
			t.Errorf("want %v; but got %v\n", 2, len(buf))
		}
	})
}
