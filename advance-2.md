Here's a **comprehensive guide** to `defer`, multi-threading (goroutines), inter-thread communication (channels), and `sync.WaitGroup` in **Go**:

---

## **1. `defer` in Go**
`defer` is used to schedule a function call to run just before the surrounding function exits.

### **Basic Example**
```go
package main

import "fmt"

func main() {
    defer fmt.Println("This runs last")
    fmt.Println("This runs first")
}
```
**Output:**
```
This runs first
This runs last
```

### **Multiple `defer` (LIFO Order)**
```go
package main

import "fmt"

func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
```
**Output:**
```
3
2
1
```
Defer calls execute in **Last In, First Out (LIFO)** order.

### **Use Case: Closing Resources**
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Ensures the file is closed when the function exits
}
```

---

## **2. Multi-threading (Goroutines)**
Go provides **lightweight threads** called **goroutines**.

### **Creating a Goroutine**
```go
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Run in a separate thread
	time.Sleep(3 * time.Second) // Give enough time for goroutine to execute
}
```

### **Wait for Goroutines (`sync.WaitGroup`)**
Since goroutines run **asynchronously**, `sync.WaitGroup` is used to wait for them to complete.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when function finishes
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter
		go worker(i, &wg)
	}

	wg.Wait() // Block until all goroutines finish
	fmt.Println("All workers completed")
}
```
**Output:**
```
Worker 1 started
Worker 2 started
Worker 3 started
Worker 1 finished
Worker 2 finished
Worker 3 finished
All workers completed
```

---

## **3. Inter-Thread Communication (`Channels`)**
Goroutines communicate using **channels**.

### **Creating and Using a Channel**
```go
package main

import "fmt"

func main() {
	ch := make(chan string) // Create a channel

	go func() {
		ch <- "Hello from Goroutine" // Send data
	}()

	msg := <-ch // Receive data
	fmt.Println(msg)
}
```
**Output:**
```
Hello from Goroutine
```

### **Buffered Channels**
Buffered channels allow sending **multiple values** without immediate reading.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3) // Buffered channel with capacity 3

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch) // Read values
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

### **Channel Closing & Range**
```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // Close channel
	}()

	for num := range ch {
		fmt.Println(num) // Read until channel is closed
	}
}
```

---

## **4. Advanced Synchronization**
### **Using `sync.Mutex` for Thread Safety**
```go
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()   // Lock before modifying shared data
	counter++   // Critical section
	mu.Unlock() // Unlock after modification
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
```

### **Using `sync.Cond` for Signal-based Coordination**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false
var cond = sync.NewCond(&sync.Mutex{})

func waitForCondition() {
	cond.L.Lock()
	for !ready {
		cond.Wait() // Wait until signaled
	}
	fmt.Println("Condition met!")
	cond.L.Unlock()
}

func main() {
	go waitForCondition()

	time.Sleep(time.Second)

	cond.L.Lock()
	ready = true
	cond.Signal() // Wake up waiting goroutine
	cond.L.Unlock()

	time.Sleep(time.Second)
}
```

---

## **5. Combining `defer`, `goroutines`, `channels`, and `WaitGroup`**
A **complete example** demonstrating all concepts together:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	time.Sleep(time.Second)
	msg := fmt.Sprintf("Worker %d finished", id)
	ch <- msg // Send message to channel
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch) // Close channel after all workers finish
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
```

---

## **Summary**
| Concept        | Description |
|---------------|-------------|
| `defer`       | Schedules a function to run at the end of the current function (LIFO execution). |
| Goroutines    | Lightweight threads used for concurrency. |
| `sync.WaitGroup` | Helps wait for multiple goroutines to finish. |
| Channels      | Used for goroutine communication. |
| `sync.Mutex`  | Ensures safe concurrent access to shared resources. |
| `sync.Cond`   | Implements signal-based coordination. |

