package main

import (
	"crypto/sha256"
)

func HashMessage(message string) [32]byte {
	return sha256.Sum256([]byte(message))
}
