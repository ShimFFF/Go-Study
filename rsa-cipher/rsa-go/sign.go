package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func SignMessage(privateKey *rsa.PrivateKey, message string) []byte {
	hashed := HashMessage(message)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		log.Fatalf("메시지 서명 실패: %v", err)
	}

	return signature
}
