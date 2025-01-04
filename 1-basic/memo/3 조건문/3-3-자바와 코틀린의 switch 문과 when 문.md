### **Java와 Kotlin에서 `switch`와 `when`문의 `break` 동작**

Java와 Kotlin의 `switch`/`when`에서 `break`가 내장되어 있는지 여부는 언어의 설계와 버전에 따라 다릅니다. 아래에서 각각의 동작을 정리

---

## **1. Java의 `switch`**

### **1-1. Java 13 이전**

- **`break`는 내장되어 있지 않음**
    - `switch`의 각 `case`가 실행된 후, 명시적으로 `break`를 작성하지 않으면 다음 `case`도 실행됨
    - 이를 **fallthrough**라고 함

### **예제**

```java

public class Main {
    public static void main(String[] args) {
        int x = 2;

        switch (x) {
            case 1:
                System.out.println("x is 1");
            case 2:
                System.out.println("x is 2");
            case 3:
                System.out.println("x is 3");
            default:
                System.out.println("x is unknown");
        }
    }
}

```

### **출력**

```csharp

x is 2
x is 3
x is unknown

```

- `case 2` 이후에 `case 3`과 `default`도 실행됨
- 이를 방지하려면 각 `case` 블록에 `break`를 추가해야 함

```java

switch (x) {
    case 1:
        System.out.println("x is 1");
        break;
    case 2:
        System.out.println("x is 2");
        break;
    case 3:
        System.out.println("x is 3");
        break;
    default:
        System.out.println("x is unknown");
}

```

---

### **1-2. Java 14 이후 (`switch` 표현식)**

Java 14부터는 `switch`를 표현식으로 사용할 수 있음

- 이 경우, **`break`가 내장되어 자동으로 동작함**.
- `case`에서 값을 반환하는 경우 `>`를 사용하며, 명시적으로 `break`를 쓸 필요가 없음.

### **예제**

```java

public class Main {
    public static void main(String[] args) {
        int x = 2;

        String result = switch (x) {
            case 1 -> "x is 1";
            case 2 -> "x is 2";
            case 3 -> "x is 3";
            default -> "x is unknown";
        };

        System.out.println(result); // 출력: x is 2
    }
}

```

---

## **2. Kotlin의 `when`**

- Kotlin의 `when`은 **항상 `break`가 내장되어 있음.**
- 각 조건 분기가 실행된 후, 자동으로 종료되며 다음 조건으로 넘어가지 않음.
- Java의 `fallthrough`와 같은 동작이 없으며, 명시적으로 `break`를 작성할 필요도 없음.

---

### **Kotlin 예제**

```kotlin

fun main() {
    val x = 2

    when (x) {
        1 -> println("x is 1")
        2 -> println("x is 2")
        3 -> println("x is 3")
        else -> println("x is unknown")
    }
}

```

### **출력**

```csharp

x is 2

```

- `when`은 자동으로 조건에 맞는 분기를 실행하고 종료함.
- 다음 조건으로 넘어가지 않음.

---

### **Kotlin에서는 `fallthrough`가 지원되지 않음**

- Kotlin은 `fallthrough`를 명시적으로 지원하지 않음.
- 여러 조건을 묶으려면 쉼표 `,`로 조건을 나열할 수 있음.

### **다중 조건 예제**

```kotlin

fun main() {
    val day = "Saturday"

    when (day) {
        "Monday", "Tuesday", "Wednesday", "Thursday", "Friday" -> println("It's a weekday")
        "Saturday", "Sunday" -> println("It's the weekend")
        else -> println("Invalid day")
    }
}

```

### **출력**

```rust

It's the weekend

```

---

## **3. Java와 Kotlin 비교**

| **특징** | **Java (13 이전)** | **Java (14 이상)** | **Kotlin** |
| --- | --- | --- | --- |
| **`break` 기본 동작** | 없음 (명시적으로 작성해야 함) | 내장되어 자동으로 다음 분기로 넘어가지 않음 | 내장되어 자동으로 다음 분기로 넘어가지 않음 |
| **`fallthrough` 동작** | 기본 동작 (명시적으로 `break` 추가해야 방지 가능) | 지원하지 않음.
기존의 `switch` 문을 사용하면 `fallthrough` 동작 가능 | 지원하지 않음 |
| **다중 조건 처리** | `case A: case B:`와 같은 방식으로 처리 | `case A, B ->`로 간결하게 처리 | 쉼표 `,`로 간결하게 처리 |
| **표현식으로 사용 가능 여부** | 불가능 | 가능 | 가능 |

---

### **결론**

1. **Java 13 이전**:
    - `switch`에서 `break`를 명시적으로 작성해야 함
    - 자동 `fallthrough` 동작 때문에 실수할 가능성이 있음
2. **Java 14 이상**:
    - `switch` 표현식이 도입되면서 `break`가 내장되어 더 간결하고 안전하게 사용할 수 있음
3. **Kotlin**:
    - `when`은 `break`가 항상 내장되어 있어 안전하며, Java의 `fallthrough` 동작을 지원하지 않음
    - `when`은 표현식과 표현문 모두로 사용할 수 있어 더욱 유연함