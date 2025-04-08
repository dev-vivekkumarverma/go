

## 🧾 1. **Data Types & Sizes**

| Category       | Data Type       | Size (Bytes)      | Notes |
|----------------|------------------|--------------------|-------|
| Boolean        | `bool`           | 1                  | `true` or `false` |
| Integer        | `int8`           | 1                  | Range: -128 to 127 |
|                | `uint8` / `byte` | 1                  | Range: 0 to 255 |
|                | `int16`          | 2                  | -32K to 32K |
|                | `uint16`         | 2                  | 0 to 65K |
|                | `int32` / `rune` | 4                  | Unicode code point |
|                | `uint32`         | 4                  | 0 to 4B |
|                | `int64`          | 8                  | Huge integers |
|                | `uint64`         | 8                  | Unsigned 64-bit |
|                | `int` / `uint`   | 4 or 8 (arch)       | 32-bit or 64-bit based |
| Floating Point | `float32`        | 4                  | IEEE 754 |
|                | `float64`        | 8                  | Default choice for float |
| Complex        | `complex64`      | 8 (2 × float32)    | Real + Imaginary |
|                | `complex128`     | 16 (2 × float64)   | High precision |
| Strings        | `string`         | 16 (header)        | 8 bytes ptr + 8 bytes len |
| Arrays         | `[n]T`           | n × size of T      | Fixed-size block |
| Slices         | `[]T`            | 24 (header only)   | ptr + len + cap |
| Pointers       | `*T`             | 4 or 8 (arch)       | Memory address |
| Interfaces     | `interface{}`    | 16                 | Type ptr + Data ptr |
| Maps           | `map[K]V`        | 8+ (header)        | Varies a lot |
| Channels       | `chan T`         | 8+ (header)        | Varies based on T |
| Functions      | `func`           | 8                  | Function pointer |

---

## 📐 2. **Memory Layout of a Struct**

In Go, struct fields are **aligned** to their type's size. This means there may be **padding** between fields.

### ⚠️ Misaligned Struct
```go
type BadStruct struct {
	A byte    // 1 byte
	B int32   // 4 bytes
	C byte    // 1 byte
}
```

Memory layout (with padding):
```
A [1] + [3 padding]
B [4]
C [1] + [3 padding]
= 12 bytes
```

### ✅ Optimized Struct
```go
type GoodStruct struct {
	B int32
	A byte
	C byte
}
```
Now:
```
B [4]
A [1]
C [1] + [2 padding]
= 8 bytes
```

👉 Use `unsafe.Sizeof()` to measure struct size and `alignas` struct tags to control layout (rarely needed).

---

## 🧠 3. **Visual Memory Layout Diagram**

```
string s = "Hi"
+------------------+
| Pointer (8 bytes)| → "H" "i"
| Length  (8 bytes)| = 2
+------------------+

[]int slice = []int{1,2,3}
+------------------+
| Pointer (8 bytes)| → array on heap
| Length  (8 bytes)| = 3
| Capacity(8 bytes)| = 3
+------------------+

interface{} v = 123
+------------------+
| Type Info (8B)   | → points to int type
| Data Ptr (8B)    | → actual value
+------------------+
```

---

## 🔍 4. **Check Sizes Dynamically in Code**

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		b  bool
		i  int
		f  float64
		r  rune
		s  string
		sl []int
	)

	fmt.Println("bool:", unsafe.Sizeof(b))
	fmt.Println("int:", unsafe.Sizeof(i))
	fmt.Println("float64:", unsafe.Sizeof(f))
	fmt.Println("rune:", unsafe.Sizeof(r))
	fmt.Println("string:", unsafe.Sizeof(s))   // Header only
	fmt.Println("[]int:", unsafe.Sizeof(sl))   // Header only
}
```

---

## 📎 5. Tips

- Use `int` unless you need a specific width (e.g., when working with files or protocols).
- Strings and slices store **headers** (pointer + length [+ capacity for slices]).
- Always consider padding when optimizing memory in `struct`s.
- Use `unsafe.Sizeof()` for inspection, not for general program logic.

