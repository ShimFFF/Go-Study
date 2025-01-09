### **`fmt.Scan` 함수 설명**

```go
func Scan(a ...any) (n int, err error)
```

### **매개변수**

- `a ...any`: 입력받은 데이터를 저장할 변수의 주소를 전달해야 함
    - 변수의 주소는 `&` 연산자를 사용해서 전달함
    - 예: `fmt.Scan(&x)` (변수 `x`에 입력 값을 저장)

### **반환 값**

1. **`n int`**: 성공적으로 읽은 값의 개수를 반환함.
2. **`err error`**: 입력 과정에서 발생한 에러를 반환. 에러가 없으면 `nil`을 반환.

---

### **사용 예제**

### 1. 숫자 입력받기

```go
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Enter a number: ")
	fmt.Scan(&x) // 입력값을 x 변수에 저장
	fmt.Println("You entered:", x)

```

### 실행 결과:

```yaml
Enter a number: 42
You entered: 42
```

---

### 2. 여러 값 입력받기

```go
package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y) // 두 개의 값을 입력받아 x와 y에 저장
	fmt.Println("You entered:", x, y)
}
```

### 실행 결과:

```yaml
Enter two numbers: 10 20
You entered: 10 20
```

---

### 3. 문자열 입력받기

```go
package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // 입력값을 name 변수에 저장
	fmt.Println("Hello,", name)
}
```

### 실행 결과:

```mathematica
Enter your name: Alice
Hello, Alice
```

---

### 4. 입력 에러 처리

입력 값이 잘못되었거나 변환할 수 없을 경우, `err`를 확인하여 처리 함.

```go
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Enter a number: ")
	n, err := fmt.Scan(&x) // 반환값 n과 err를 확인
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully read", n, "value(s):", x)
	}
}
```

### 실행 결과:

```vbnet
Enter a number: abc
Error: expected integer
```

---

### **주의 사항**

1. **공백 구분**
    - `fmt.Scan`은 입력값을 공백(스페이스, 탭, 줄바꿈 등)으로 구분함.
    - 여러 값을 입력할 때 공백으로 구분해야 함.
2. **EOF 처리**
    - 입력이 끝났거나 예상보다 적은 값이 입력되었을 경우, `Scan`은 읽은 값의 개수를 반환하고 에러를 보고함.
3. **주소 전달**
    - 변수의 값을 변경하려면 `&` 연산자를 사용해 변수의 주소를 전달해야 함.
    - 예: `fmt.Scan(&x)`
4. **에러 처리**
    - 항상 `err`를 확인하여 입력이 성공적으로 처리되었는지 확인하는 것이 좋음.

---

### **프로젝트 코드에서의 사용**

```go
fmt.Scan(&choice) // choice 변수에 사용자가 입력한 값을 저장
```

### 동작:

1. 사용자가 입력한 값을 `choice` 변수에 저장.
2. 만약 사용자가 숫자가 아닌 값을 입력하거나, 입력이 잘못된 경우 에러가 발생할 수 있음.

### 입력 예:

```mathematica
Choose a shape (1: Circle, 2: Rectangle, 3: Exit): 1
```

- 입력값 `1`이 `choice` 변수에 저장됨.

---

`fmt.Scan`은 간단한 입력 처리를 위한 유용한 함수임. 더 복잡한 입력 처리를 위해서는 `fmt.Scanln`, `fmt.Sscanf`, 또는 `bufio.Scanner`와 같은 다른 입력 방법을 사용할 수도 있음