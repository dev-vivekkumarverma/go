# Go Programming Guide

## 1. Variable Scope

### Local Variables
Local variables are defined inside a function and are accessible only within that function. They help in encapsulating data and preventing unintended modifications by other parts of the code.

- Declared using `var` or `:=` (short declaration syntax).

```go
package main
import "fmt"

func main() {
    var message string = "Hello, Go!"  // Using var
    count := 10  // Using := (short declaration)
    fmt.Println(message, count)
}
```

### Package-Level Variables
Package-level variables are accessible only within the package. They help in sharing data between multiple functions in the same package.

- Declared at the beginning, after the `import` statement.

```go
package main
import "fmt"

var packageVar = "Accessible within the package"

func main() {
    fmt.Println(packageVar)
}
```

### Global Variables and Functions
Global variables and functions are accessible across the entire application. They are useful for storing configurations and reusable utility functions.

- Must be exported by capitalizing their names.

```go
package main
import "fmt"

var GlobalVar = "I am global"

func GlobalFunction() {
    fmt.Println("I am a global function")
}
```

## 2. Importing Functions from Other Modules
Modules allow code reuse and separation of concerns. You can import functions from other modules as follows:

```go
import "myapp/utility"
```

- The function needs to be exported by capitalizing its name in the module where it is defined.

## 3. Data Types in Go
Go provides a rich set of data types to handle various types of data efficiently.

### Integer Types

- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `uintptr` for storing pointer addresses

```go
var a int = 10
var b int8 = 127
```

### Float Types

Floating-point numbers are used for fractional calculations.

- `float32`, `float64`

```go
var pi float64 = 3.1415
```

### Boolean Type
Booleans are used for conditional statements and logical operations.

```go
var isAvailable bool = true
```

### String Type
Strings store text data and are immutable.

```go
var message string = "Hello, Go!"
```

### Complex Number Types
Go supports complex numbers with real and imaginary parts.

- `complex64`, `complex128`

```go
var c complex64 = complex(2, 3)
```

### Array and Slice Types
Arrays and slices store collections of elements.

```go
var arr [3]int = [3]int{1, 2, 3} // Fixed size array
slice := []int{1, 2, 3, 4, 5}  // Dynamic size
```

### Map Type (Key-Value Pair)
Maps are used for fast lookup operations.

```go
var myMap map[string]int = map[string]int{"a": 1, "b": 2}
```

## 4. Operations in Go

### Bitwise Operations
Useful for performance optimization, cryptography, and low-level programming.

```go
and := 5 & 3  // AND operation
or := 5 | 3   // OR operation
xor := 5 ^ 3  // XOR operation
leftShift := 5 << 1  // Left shift
rightShift := 5 >> 1 // Right shift
```

## 5. Loops and Control Flow

### Range Loop
Used to iterate over slices, maps, and arrays.

```go
nums := []int{1, 2, 3, 4}
for index, value := range nums {
    fmt.Println(index, value)
}
```

## 6. Functions

### Anonymous Functions
Useful for short, one-time-use functions.

```go
func() {
    fmt.Println("Hello from an anonymous function!")
}()
```

### Variadic Functions
Accept a variable number of arguments.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## 7. Object-Oriented Concepts

### Embedding (Instead of Inheritance)
Go uses composition instead of classical inheritance.

```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}
```

## 8. Goroutines and Concurrency

### Goroutine
Lightweight thread for parallel execution.

```go
go func() {
    fmt.Println("Running in a goroutine")
}()
```

### Channels
Used for communication between goroutines.

```go
messages := make(chan string)
go func() { messages <- "Hello, Channel!" }()
fmt.Println(<-messages)
```

## 9. File Handling

### Reading Large Files with Buffers
Efficient way to handle large files without consuming too much memory.

```go
import (
    "bufio"
    "os"
)

file, _ := os.Open("largefile.txt")
defer file.Close()
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

## 10. Error Handling

### Panic and Recover
Used for handling unexpected errors gracefully.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}()
panic("Something went wrong!")
```

## 11. JSON Handling

Used for data serialization and API communication.

```go
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

data := `{"name":"John", "age":30}`
var person Person
json.Unmarshal([]byte(data), &person)
fmt.Println(person.Name, person.Age)
```

## 12. Advanced Topics

### Interfaces and Polymorphism
Go uses interfaces for polymorphism.

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }
```

### Dependency Management
Use `go mod` for package management.

```sh
go mod init myapp
go get github.com/gin-gonic/gin
```

### Unit Testing
Write tests with the `testing` package.

```go
import "testing"
func TestSum(t *testing.T) {
    result := sum(1, 2, 3)
    if result != 6 {
        t.Errorf("Expected 6 but got %d", result)
    }
}
```


