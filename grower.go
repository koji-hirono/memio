package memio

import (
	"errors"
)

type Grower interface {
	Grow(int) error
	Bytes() []byte
}

var ErrNoMem = errors.New("out of memory")
