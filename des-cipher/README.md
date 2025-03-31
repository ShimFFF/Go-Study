# DES 구현

1. 입출력 타입  
블록 크기: 64비트 (8바이트)  
입력: ASCII 문자열 "ABCDEFGH" → []byte  
키: 0x133457799BBCDFF1 → []byte 또는 uint64  

## DES 알고리즘 요약
DES는 다음과 같은 16-라운드 Feistel 구조를 따름

1. 64비트 입력 → 초기 순열 (IP) 적용
2. 32비트 L, 32비트 R로 분리
3. 16회 라운드 반복 
```
Li = Ri-1 
Ri = Li-1 XOR f(Ri-1, Ki)
```

4. 마지막에 L, R을 Swap → 최종 순열 (FP) 적용

**Feistel 함수 f**
- 32비트 입력 R → E-확장 (32 → 48비트)
- 48비트 서브키와 XOR  
- 8개의 S-box로 나눠 6비트씩 → 4비트로 압축 → 총 32비트
- P-box 순열 적용  

**Key Schedule**
- 64비트 키 → PC-1로 56비트 압축
- 두 부분으로 나눠 좌우 shift → PC-2로 48비트 서브키 16개 생성

