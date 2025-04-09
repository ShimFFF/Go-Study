# RSA 기반 디지털 서명 Go 구현

## 목적
RSA 알고리즘을 사용하여 디지털 서명 프로세스를 구현합니다. 이를 통해 디지털 서명의 필요성과 원리를 이해하고, 메시지 변조 시 검증 단계에서 탐지할 수 있음을 확인합니다.

## 구현 사항
1. **RSA 키 생성**
   - 공개 키와 개인 키를 생성합니다.
   - 키 길이: 2048 bit

2. **메시지 서명**
   - 메시지의 해시(SHA-256)를 계산하고, 개인 키로 서명합니다.
   - Padding: PSS (Probabilistic Signature Scheme)

3. **서명 검증**
   - 원본 메시지: 공개 키를 사용해 서명 검증
   - 변조된 메시지 (Eve 가 Bob 계좌를 자신의 계좌로 변경): 서명 불일치 확인

## 메시지
- 원본: "Transfer $100 from Alice's account (123-4567-1111-9999) to Bob's account (123-3456-2222-8888)."
- 변조: "Transfer $100 from Alice's account (123-4567-1111-9999) to Bob's account (123-6789-3333-7777)."

## 실행 방법
```bash
go run main.go
```