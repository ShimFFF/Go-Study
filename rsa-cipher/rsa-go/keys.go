package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("키 생성 실패: %v", err)
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}
