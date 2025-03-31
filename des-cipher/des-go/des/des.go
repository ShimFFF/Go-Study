package des

// DES 핵심 Round 함수
// 32비트 입력 + 48비트 서브키 → 32비트 출력
func feistelFunction(right uint32, subkey uint64) uint32 {
	// E-box 확장: 32비트 → 48비트
	expanded := Permute32to48(right, expansionTable)

	// 서브키와 XOR
	xored := expanded ^ subkey

	// S-box 처리
	var sboxOutput uint32 = 0
	for i := 0; i < 8; i++ {
		// 오른쪽부터 6비트씩 추출
		sixBits := (xored >> uint((7-i)*6)) & 0x3F

		// S-box 인덱스 계산
		// row: 첫 번째와 마지막 비트 (B1B6)
		// col: 중간 4비트 (B2B3B4B5)
		row := ((sixBits & 0x20) >> 4) | (sixBits & 0x01)
		col := (sixBits >> 1) & 0x0F

		// S-box 조회
		sboxVal := sBoxes[i][row][col]

		// 4비트씩 이어붙이기
		sboxOutput = (sboxOutput << 4) | uint32(sboxVal)
	}

	// P-box 순열
	return Permute32(sboxOutput, pBox)
}

// 암호화
func EncryptBlock(block uint64, roundKeys [16]uint64) uint64 {
	block = InitialPermutation(block)

	L := uint32(block >> 32)
	R := uint32(block & 0xFFFFFFFF)

	for i := 0; i < 16; i++ {
		temp := L
		L = R
		R = temp ^ feistelFunction(R, roundKeys[i])
	}

	preOutput := (uint64(R) << 32) | uint64(L)
	return FinalPermutation(preOutput)
}

// 복호화
func DecryptBlock(block uint64, roundKeys [16]uint64) uint64 {
	block = InitialPermutation(block)

	L := uint32(block >> 32)
	R := uint32(block & 0xFFFFFFFF)

	for i := 15; i >= 0; i-- {
		temp := R
		R = L ^ feistelFunction(R, roundKeys[i])
		L = temp
	}

	preOutput := (uint64(R) << 32) | uint64(L)
	return FinalPermutation(preOutput)
}
