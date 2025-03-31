package des

// 키 스케줄 구현
func GenerateRoundKeysFromBytes(rawKey []byte) [16]uint64 {
	var roundKeys [16]uint64
	if len(rawKey) != 8 {
		panic("DES key must be 8 bytes")
	}

	// 정확한 bit 위치로 순열 적용
	permutedKey := PermuteBitsFromBytes(rawKey, pc1)

	C := uint32((permutedKey >> 28) & 0x0FFFFFFF)
	D := uint32(permutedKey & 0x0FFFFFFF)

	for i := 0; i < 16; i++ {
		C = leftShift28(C, shifts[i])
		D = leftShift28(D, shifts[i])
		combined := (uint64(C) << 28) | uint64(D)

		// PC2도 동일하게 정확한 bit 위치로 적용
		combinedBytes := make([]byte, 7)
		for j := 0; j < 7; j++ {
			combinedBytes[6-j] = byte((combined >> (j * 8)) & 0xFF)
		}
		roundKeys[i] = PermuteBitsFromBytes(combinedBytes, pc2)
	}
	return roundKeys
}

// 28비트 좌측 순환 시프트
func leftShift28(val uint32, n int) uint32 {
	return ((val << n) | (val >> (28 - n))) & 0x0FFFFFFF
}
