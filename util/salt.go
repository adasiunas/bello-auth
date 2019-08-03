package util

import (
	"crypto/rand"
	"encoding/hex"
)

type Salt []byte

func GenerateSalt() Salt {
	var salt = make([]byte, 32)
	rand.Read(salt)
	return Salt(salt)
}

func (s Salt) String() string {
	var salt = make([]byte, 64)
	hex.Encode(salt, s)
	return string(salt)
}

func FromStringToSalt(hexSalt string) Salt {
	var salt = make([]byte, 32)
	hex.Decode(salt, []byte(hexSalt))
	return Salt(salt)
}
