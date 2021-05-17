package memio

import (
	"testing"
)

func TestVar_Grow(t *testing.T) {
	t.Run("0 -> 2byte", func(t *testing.T) {
		m := NewVar(nil)
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
		if cap(buf) != 64 {
			t.Errorf("want %v; but got %v\n", 64, cap(buf))
		}
		if len(buf) != 2 {
			t.Errorf("want %v; but got %v\n", 2, len(buf))
		}
	})
	t.Run("2 -> 2byte", func(t *testing.T) {
		p := make([]byte, 2)
		m := NewVar(p)
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
	t.Run("2 -> 4byte", func(t *testing.T) {
		p := make([]byte, 2)
		m := NewVar(p)
		err := m.Grow(3)
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
		if cap(buf) != 4 {
			t.Errorf("want %v; but got %v\n", 4, cap(buf))
		}
		if len(buf) != 3 {
			t.Errorf("want %v; but got %v\n", 3, len(buf))
		}
	})
}
