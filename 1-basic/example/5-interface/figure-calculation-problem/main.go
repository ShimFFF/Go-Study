package main

import (
	"fmt"
	"project/shape"
)

func main() {
	var choice int
	for {
		fmt.Println("\nChoose a shape (1: Circle, 2: Rectangle, 3: Exit):")
		fmt.Scan(&choice)

		var figure shape.Shape

		switch choice {
		case 1: // Circle 선택
			var radius float64
			fmt.Print("Enter radius: ")
			fmt.Scan(&radius)

			// Circle 생성 및 출력
			figure = shape.NewCircle(radius)
			// := -> 단축 변수 선언 및 초기화 (새로운 변수 만들 때)
			// = -> 변수 할당 (기존 변수에 값을 넣을 때)

		case 2: // Rectangle 선택
			var width, height float64
			fmt.Print("Enter width: ")
			fmt.Scan(&width)
			fmt.Print("Enter height: ")
			fmt.Scan(&height)

			// Rectangle 생성 및 출력
			figure = shape.NewRectangle(width, height)

		case 3: // 프로그램 종료
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		// 추상화된 Shape 인터페이스를 사용하여 출력
		fmt.Printf("Area: %.2f\n", figure.Area())
		fmt.Printf("Perimeter: %.2f\n", figure.Perimeter())
	}
}
