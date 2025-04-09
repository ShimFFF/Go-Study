package main

import (
	"fmt"
)

func main() {
	// 키 생성
	privateKey, publicKey := GenerateKeys()

	// 메시지들
	originalMessage := "Transfer $100 from Alice's account (123-4567-1111-9999) to Bob's account (123-3456-2222-8888)."
	tamperedMessage := "Transfer $100 from Alice's account (123-4567-1111-9999) to Bob's account (123-6789-3333-7777)."

	// 메시지 서명
	signature := SignMessage(privateKey, originalMessage)

	fmt.Println("RSA 키 생성 완료")
	fmt.Println("원본 메시지 서명 완료")
	fmt.Printf("Signature: %x\n\n", signature)

	// 서명 검증
	fmt.Println("원본 메시지 검증:")
	VerifySignature(publicKey, originalMessage, signature)

	fmt.Println("\n변조 메시지 검증:")
	VerifySignature(publicKey, tamperedMessage, signature)
}
