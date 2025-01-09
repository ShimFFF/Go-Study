package shape

import "math"

// Circle 구조체 정의
type Circle struct {
	radius float64 // 프라이빗 필드
}

// NewCircle: Circle 생성자 함수
func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

// Getter: Circle의 Radius 반환
func (c *Circle) Radius() float64 {
	return c.radius
}

// Setter: Circle의 Radius 설정
func (c *Circle) SetRadius(radius float64) {
	if radius > 0 {
		c.radius = radius
	}
}

// Area 메서드: Circle의 면적 계산
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter 메서드: Circle의 둘레 계산
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
