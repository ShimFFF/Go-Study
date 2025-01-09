package shape

// Shape 인터페이스 정의
type Shape interface {
	Area() float64      // 면적 계산
	Perimeter() float64 // 둘레 계산
}
