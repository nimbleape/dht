package bep44

import (
	"crypto/sha256"
)

type Target = [sha256.Size]byte

func MakeMutableTarget(pubKey [32]byte, salt []byte) Target {
	return sha256.Sum256(append(pubKey[:], salt...))
}
