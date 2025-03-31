package des

import "encoding/binary"

// 바이트 배열을 uint64로 변환 (Big Endian)
func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// uint64를 바이트 배열로 변환 (Big Endian)
func Uint64ToBytes(n uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	return buf
}
