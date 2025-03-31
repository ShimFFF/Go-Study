package main

import (
	"descipher/des"
	"fmt"
	"time"
)

func main() {
	plaintext := []byte("ABCDEFGH") // 정확히 8바이트
	if len(plaintext) != 8 {
		panic("Plaintext must be 8 bytes")
	}

	// 키 및 라운드 키 생성
	keyBytes := []byte{0x13, 0x34, 0x57, 0x79, 0x9B, 0xBC, 0xDF, 0xF1}
	roundKeys := des.GenerateRoundKeysFromBytes(keyBytes)

	// 평문 블록 변환
	block := des.BytesToUint64(plaintext)
	fmt.Printf("Original plaintext (hex): %016X\n", block)

	// 암호화
	cipher := des.EncryptBlock(block, roundKeys)
	fmt.Printf("Encrypted ciphertext (hex): %016X\n", cipher)

	// 복호화
	decrypted := des.DecryptBlock(cipher, roundKeys)
	recovered := des.Uint64ToBytes(decrypted)
	fmt.Printf("Decrypted result (hex): %016X\n", decrypted)
	fmt.Printf("Recovered text: %s\n", string(recovered))

	// 성능 측정: 100번 반복
	startEnc := time.Now()
	for i := 0; i < 100; i++ {
		_ = des.EncryptBlock(block, roundKeys)
	}
	elapsedEnc := time.Since(startEnc)

	startDec := time.Now()
	for i := 0; i < 100; i++ {
		_ = des.DecryptBlock(cipher, roundKeys)
	}
	elapsedDec := time.Since(startDec)

	fmt.Printf("\nEncryption 100x time: %s\n", elapsedEnc)
	fmt.Printf("Decryption 100x time: %s\n", elapsedDec)
}
