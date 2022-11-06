package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenID() string {
	b := make([]byte, 5) //equals 10 characters
	rand.Read(b)
	s := hex.EncodeToString(b)
	return s
}
