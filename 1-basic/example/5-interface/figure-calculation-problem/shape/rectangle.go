package shape

// Rectangle 구조체 정의
type Rectangle struct {
	width  float64 // 프라이빗 필드
	height float64 // 프라이빗 필드
}

// NewRectangle: Rectangle 생성자 함수
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

// Getter: Rectangle의 Width 반환
func (r *Rectangle) Width() float64 {
	return r.width
}

// Setter: Rectangle의 Width 설정
func (r *Rectangle) SetWidth(width float64) {
	if width > 0 {
		r.width = width
	}
}

// Getter: Rectangle의 Height 반환
func (r *Rectangle) Height() float64 {
	return r.height
}

// Setter: Rectangle의 Height 설정
func (r *Rectangle) SetHeight(height float64) {
	if height > 0 {
		r.height = height
	}
}

// Area 메서드: Rectangle의 면적 계산
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter 메서드: Rectangle의 둘레 계산
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
