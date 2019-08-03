package util

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

type Hash []byte

func HashFromString(hashString string) Hash {
	bytes := make([]byte, 128)
	hex.Decode(bytes, []byte(hashString))
	return Hash(bytes)
}

func GetHash(str string, salt Salt) Hash {
	hash := pbkdf2.Key([]byte(str), salt, 65536, 128, sha256.New)
	return Hash(hash)
}

func (hash Hash) String() string {
	hashHex := make([]byte, 256)
	hex.Encode(hashHex, []byte(hash))
	return string(hashHex)
}
