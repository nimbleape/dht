package dht

import (
	"crypto/sha256"
)

func HashTuple(bs ...[]byte) (ret [32]byte) {
	h := sha256.New()
	for _, b := range bs {
		h.Reset()
		h.Write(ret[:])
		h.Write(b)
		h.Sum(ret[:0])
	}
	return
}
