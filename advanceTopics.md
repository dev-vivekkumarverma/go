# Advanced Go Topics

## 1. Interfaces and Polymorphism
Go uses interfaces for achieving polymorphism, allowing different types to implement the same behavior without explicit inheritance.

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

func main() {
    var s Speaker
    s = Dog{}
    fmt.Println(s.Speak()) // Output: Woof!
    s = Cat{}
    fmt.Println(s.Speak()) // Output: Meow!
}
```

## 2. Dependency Management
Go modules (`go mod`) help manage dependencies efficiently.

```sh
go mod init myapp
go get github.com/gin-gonic/gin
```

## 3. Unit Testing
Testing in Go is performed using the `testing` package.

```go
import "testing"
func TestSum(t *testing.T) {
    result := sum(1, 2, 3)
    if result != 6 {
        t.Errorf("Expected 6 but got %d", result)
    }
}
```

## 4. Logging and Debugging
Use the `log` package for logging messages and errors.

```go
import "log"
func main() {
    log.Println("This is a log message")
}
```

## 5. Networking in Go
### HTTP Server
```go
import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### HTTP Client
```go
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, _ := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

## 6. Database Interaction
### Connecting to PostgreSQL
```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    connStr := "user=username dbname=mydb sslmode=disable"
    db, _ := sql.Open("postgres", connStr)
    defer db.Close()
}
```

---
## Advance Level-2
- covers `core concurrency patterns` in Go. 
- <a href="./advance-2.md"><i>Click here</i></a>