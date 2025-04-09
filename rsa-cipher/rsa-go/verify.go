package main

import (
	"crypto"
	"crypto/rsa"
	"fmt"
)

func VerifySignature(publicKey *rsa.PublicKey, message string, signature []byte) {
	hashed := HashMessage(message)

	err := rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		fmt.Println("서명 검증 실패:", err)
	} else {
		fmt.Println("서명 검증 성공!")
	}
}
