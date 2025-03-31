package des

// 초기 순열
func InitialPermutation(input uint64) uint64 {
	return Permute(input, initialPermutation)
}

// 최종 순열
func FinalPermutation(input uint64) uint64 {
	return Permute(input, finalPermutation)
}

// 공통 순열 함수 (64비트 입력에 테이블 적용)
func Permute(input uint64, table []int) uint64 {
	return PermuteN(input, table, len(table))
}

// 32비트 입력 → 48비트로 확장 (E-box 순열에 사용)
func Permute32to48(input uint32, table []int) uint64 {
	var output uint64 = 0
	for i, pos := range table {
		bit := (input >> (32 - pos)) & 1
		output |= uint64(bit) << (47 - i)
	}
	return output
}

// 바이트 배열 기준으로 비트 순열 (PC1, PC2 전용)
func PermuteBitsFromBytes(input []byte, table []int) uint64 {
	var output uint64 = 0
	for i, pos := range table {
		byteIndex := (pos - 1) / 8
		bitIndex := 7 - ((pos - 1) % 8) // MSB 기준

		bit := (input[byteIndex] >> bitIndex) & 1
		output |= uint64(bit) << (len(table) - 1 - i)
	}
	return output
}

// 32비트 입력 순열 (P-box에 사용)
func Permute32(input uint32, table []int) uint32 {
	var output uint32 = 0
	for i, pos := range table {
		bit := (input >> (32 - pos)) & 1
		output |= bit << (31 - i)
	}
	return output
}

// 일반화된 순열 함수: 입력 비트에서 테이블 위치대로 outLen 크기만큼 추출
func PermuteN(input uint64, table []int, outLen int) uint64 {
	var output uint64 = 0
	for i, pos := range table {
		bit := (input >> (64 - pos)) & 1
		output |= bit << (outLen - 1 - i)
	}
	return output
}
