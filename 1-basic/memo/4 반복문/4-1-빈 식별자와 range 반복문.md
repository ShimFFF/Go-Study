---

## **`range`의 동작**

`for ... range`는 **슬라이스(Slice)**, **배열(Array)**, **맵(Map)**, **문자열(String)**, 또는 **채널(Channel)**에서 요소를 순회(iterate)하기 위해 사용됨.

### **구조**

```go

for index, value := range collection {
    // 반복 블록
}

```

- **`collection`**: 순회할 대상 (슬라이스, 배열, 맵 등).
- **`index`**: 현재 반복의 인덱스(또는 키).
- **`value`**: 현재 인덱스(또는 키)에 해당하는 값.

## **2. 예제: 슬라이스에서의 사용**

### **코드**

```go

package main

import "fmt"

func main() {
    nums := []int{10, 20, 30, 40}

    for i, v := range nums {
        fmt.Printf("Index %d: %d\n", i, v)
    }
}

```

### **출력**

```yaml

Index 0: 10
Index 1: 20
Index 2: 30
Index 3: 40

```

### **해석**

1. **`nums`**:
    - 슬라이스 `[10, 20, 30, 40]`를 순회.
2. **`i`**:
    - 현재 요소의 인덱스 (`0`, `1`, `2`, `3`).
3. **`v`**:
    - 현재 인덱스에 해당하는 값 (`10`, `20`, `30`, `40`).

---

## **3. 슬라이스와 배열에서의 `for ... range` 동작**

- **슬라이스**:
    - 슬라이스를 `for ... range`로 순회하면, 순서대로 인덱스와 값을 반환합니다.
- **배열**:
    - 배열도 슬라이스와 동일하게 동작.

### **예제: 배열**

```go

package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}

    for i, v := range arr {
        fmt.Printf("Index %d: %d\n", i, v)
    }
}

```

### **출력**

```yaml

Index 0: 1
Index 1: 2
Index 2: 3

```

---

## **4. 맵(Map)에서의 `for ... range` 동작**

맵에서 `for ... range`는 **키-값 쌍**을 순회합니다.

### **예제: 맵 순회**

```go

package main

import "fmt"

func main() {
    ages := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Charlie": 35,
    }

    for key, value := range ages {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}

```

### **출력**

```yaml

Key: Alice, Value: 25
Key: Bob, Value: 30
Key: Charlie, Value: 35

```

### **특징**

- **`key`**: 맵의 각 키.
- **`value`**: 키에 해당하는 값.

---

## **5. 문자열(String)에서의 `for ... range` 동작**

문자열은 UTF-8 인코딩된 문자 배열로 취급되며, `for ... range`로 각 문자의 **인덱스와 유니코드 코드포인트**를 순회합니다.

### **예제: 문자열 순회**

```go

package main

import "fmt"

func main() {
    str := "Go언어"

    for i, r := range str {
        fmt.Printf("Index: %d, Rune: %c\n", i, r)
    }
}

```

### **출력**

```yaml

Index: 0, Rune: G
Index: 1, Rune: o
Index: 2, Rune: 언
Index: 5, Rune: 어

```

### **특징**

- *`r`*는 각 문자의 유니코드 값(`rune` 타입).
- 문자열은 UTF-8로 인코딩되어 있으므로, 한글과 같은 다중 바이트 문자는 인덱스 간격이 2 이상일 수 있음.

---

## **6. 무시 가능한 값**

- `for ... range`에서 특정 값을 사용하지 않을 경우, **언더스코어(`_`)**로 무시할 수 있습니다.

### **예제: 인덱스 무시**

```go

for _, value := range nums {
    fmt.Println(value) // 값만 출력
}

```

### **예제: 값 무시**

```go

for index := range nums {
    fmt.Println(index) // 인덱스만 출력
}

```

---

## **7. 주의 사항**

1. **값 복사**
    - `for ... range`는 슬라이스나 맵의 **복사본**을 순회합니다.
    - 따라서 값 타입의 슬라이스 요소를 변경해도 원본에 영향을 주지 않습니다.
2. **참조 타입**
    - 참조 타입의 슬라이스 요소를 수정하면 원본에 영향을 줍니다.

### **예제: 슬라이스 값 변경**

```go

package main

import "fmt"

func main() {
    nums := []int{1, 2, 3}

    for i, v := range nums {
        nums[i] = v * 2 // 원본 수정
    }

    fmt.Println(nums) // [2, 4, 6]
}

```

---

## **8. 요약**

| **컬렉션** | **키/인덱스** | **값** | **특징** |
| --- | --- | --- | --- |
| **슬라이스/배열** | 인덱스 (`i`) | 요소 값 (`v`) | 순서대로 인덱스와 값을 반환. |
| **맵** | 키 (`key`) | 키에 해당하는 값 (`value`) | 키-값 쌍을 임의 순서로 반환. |
| **문자열** | 인덱스 (`i`) | 유니코드 코드포인트 (`r`) | UTF-8 기반 문자 단위로 순회. |

`for i, v := range nums`는 Go에서 간결하고 강력한 컬렉션 순회 방

---

### **`_`의 의미**

- Go에서 `_`는 **빈 식별자(blank identifier)**로 사용됩니다.
- `_`는 값을 무시하고, 해당 위치에 어떤 값이 들어오더라도 사용하지 않겠다는 의미입니다.
- 이 경우, `index` 값을 무시하고 `value`만 사용하고 싶을 때 `_`를 사용합니다.

---

### **`for _, num := range numbers`**

```go

for _, num := range numbers {
    // 반복 블록
}

```

- `range numbers`는 슬라이스 `numbers`를 순회하며, 각 요소의 인덱스와 값을 반환합니다.
- `_`는 인덱스를 무시하겠다는 뜻입니다.
- `num`은 현재 요소의 값을 담는 변수입니다.

---

### **자세한 설명**

예제를 기준으로 `numbers`의 값은 `[1, 2, 3]`입니다.

`for _, num := range numbers`는 아래처럼 동작합니다:

- 첫 번째 반복: `num = 1` (인덱스 `0`은 무시)
- 두 번째 반복: `num = 2` (인덱스 `1`은 무시)
- 세 번째 반복: `num = 3` (인덱스 `2`는 무시)

이렇게 `num`만 사용하여 슬라이스의 값들을 순회하게 됩니다.

---

### **`_`를 사용하지 않은 형태**

만약 인덱스를 활용하고 싶다면 `_`를 제외하고 인덱스를 사용할 수 있습니다.

```go

package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3}

    for i, num := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", i, num)
    }
}

```

### **출력**

```yaml

Index: 0, Value: 1
Index: 1, Value: 2
Index: 2, Value: 3

```

---

### **`_` 사용 이유**

1. **필요 없는 값을 무시할 때**
    - 인덱스를 사용하지 않을 경우, `_`를 사용하여 코드를 간결하게 만듭니다.
2. **다른 값을 덮어쓰지 않음**
    - `_`는 값이 저장되지 않으므로 메모리 공간을 할당받지 않습니다.

---

### **`_`의 활용 예제**

### **인덱스만 필요할 때**

```go

package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3}

    for i, _ := range numbers {
        fmt.Println("Index:", i)
    }
}

```

### **값만 필요할 때**

```go

package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3}

    for _, num := range numbers {
        fmt.Println("Value:", num)
    }
}

```

---