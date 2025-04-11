package main

import (
	"datatype_skill/chanpractice"
	"datatype_skill/interfacehandler"
	"datatype_skill/jsonhandling"
	stringManipulation "datatype_skill/string"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

func main() {

	// interface

	fmt.Println(interfacehandler.TESTTEXT)

	r := Rectangle{l: 2.0, b: 3.0}
	c := Circle{r: 5.0}
	s := Square{a: 10.0}

	PrintDimensions(r)
	PrintDimensions(c)
	PrintDimensions(s)
}

// use of interface with struct
type Rectangle struct {
	l, b float32
}
type Circle struct {
	r float32
}
type Square struct {
	a float32
}

func (r Rectangle) Area() float64 {
	return float64(r.l) * float64(r.b)
}

func (r Rectangle) Circumference() float64 {
	return 2 * (float64(r.l) + float64(r.b))
}

func (s Square) Area() float64 {
	return float64(s.a) * float64(s.a)
}

func (s Square) Circumference() float64 {
	return 4 * float64(s.a)
}

func (c Circle) Area() float64 {
	return (22 / 7) * float64(c.r) * float64(c.r)
}

func (c Circle) Circumference() float64 {
	return 2 * (22 / 7) * float64(c.r)
}
func PrintDimensions(s interfacehandler.Shape) {
	fmt.Println("Area::", s.Area(), " sq units")
	fmt.Println("Circumference::", s.Circumference(), " units")

}

// old data struncture function calls
func oldTest() {
	fmt.Println(stringManipulation.GiveName())
	fmt.Println(stringManipulation.Concate("Vivek", " Verma"))
	fmt.Println(len(stringManipulation.Concate("Vivek", " Verma")))

	var str string = "Vivek Kumar Verma"

	for ind, r := range str {
		fmt.Printf("%T | %v | %d | %c | %U \n", r, r, ind, r, r)
	}

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	fmt.Println(strings.Compare("Vivek", "Vivek")) // 1: str1>str2; 0: str1==str2; -1:str1<str2

	str = "golang is fun"

	fmt.Println(strings.Contains(str, "lang"))
	fmt.Println(strings.HasPrefix(str, "go"))
	fmt.Println(strings.HasSuffix(str, "fun"))
	fmt.Println(strings.Index(str, "is"))
	fmt.Println("#############")
	fmt.Println(strings.Contains(str, "laxyzng"))
	fmt.Println(strings.HasPrefix(str, "fun"))
	fmt.Println(strings.HasSuffix(str, "is"))
	fmt.Println(strings.Index(str, "x"))

	// splitting and joining words

	// var str string=""
	var words []string = make([]string, 0)
	str = "string,int,bool"
	fmt.Println(string(str[11]))

	words = strings.Split(str, ",")

	fmt.Printf("%v : %T\n", words, words)

	fmt.Println(strings.Join(words, "; "))

	words = strings.Split(strings.Join(words, "; "), ";")

	new_words := make([]string, 0)

	for _, word := range words {
		fmt.Println(word, "===", len(word))
		new_words = append(new_words, strings.TrimSpace(word))
	}

	fmt.Println(new_words, "===", strings.Join(new_words, "."))
	for _, word := range new_words {
		fmt.Println(word, "==", len(word))
	}

	// using rune

	runes := []rune(str)
	fmt.Println(runes)
	fmt.Println(len(str)) // bytes: 6
	fmt.Println(len(runes))

	s := "Go世界🌏"
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}

	var xyz []rune

	xyz = append(xyz, int32(86))
	xyz = append(xyz, int32(105))
	xyz = append(xyz, int32(118))
	xyz = append(xyz, int32(101))
	xyz = append(xyz, int32(107))
	xyz = append(xyz, int32(32))
	xyz = append(xyz, int32(75))

	fmt.Printf("%v : %T\n", xyz, xyz)

	fmt.Println(string(xyz))

	// From rune (Unicode code point)
	var r rune = 0x4E16 // '世'
	fmt.Println("From rune:", string(r), r)

	// From UTF-8 bytes
	utf8Bytes := []byte{0xe4, 0xb8, 0x96}
	fmt.Println("From bytes:", string(utf8Bytes), utf8Bytes)

	b := []byte{0xe4, 0xb8, 0x96} // UTF-8 for 世
	r, size := utf8.DecodeRune(b)
	fmt.Printf("Rune: %c, Size: %d , value of rune: %v\n", r, size, r)

	// chain practice

	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(6)
	go chanpractice.SendMessage("Hey this message is a test of channel1 working.", channel1, &wg, 5)
	go chanpractice.ReceiveMessage(channel1, &wg)
	go chanpractice.SendMessage("Hey this message is a test of channel2 working.", channel2, &wg, 2)
	go chanpractice.ReceiveMessage(channel2, &wg)
	go chanpractice.SendMessage("Hey this message is a test of channel3 working.", channel3, &wg, 3)
	go chanpractice.ReceiveMessage(channel3, &wg)

	wg.Wait()

	// fan-in /fan-out
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}

	// multiprocessing using go routine

	start := time.Now()
	numCPUs := runtime.NumCPU()
	fmt.Println("🧠 CPU cores available:", numCPUs)

	// Allow Go runtime to use all available CPUs
	runtime.GOMAXPROCS(numCPUs)

	// var wg sync.WaitGroup
	for i := 1; i <= numCPUs+100; i++ {
		wg.Add(1)
		go heavyTask(i, &wg)
	}
	wg.Wait()
	fmt.Println("✅ All tasks done. time taken ::", time.Since(start))

	// json and struct
	// using marshel for converting the json to struct with some constraints

	person := jsonhandling.JsonToStruct()
	fmt.Printf("\nvalue : %v and type : %T\n", person, person)
	fmt.Println(person.Naam, person.Umar, person.Phone)
	person.Phone = "13242534"
	jsonString := jsonhandling.StructToJson(person)
	fmt.Printf("\nvalue: %v | type : %T\n", jsonString, jsonString)

	// reverse of string
	fmt.Println("reverse string of::", stringManipulation.ReveseString("Hello India"))

}

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Println("tasks ", id, "started at ::", start)
	sum := 0
	for i := 0; i < 1e9; i++ {
		sum += i
	}
	fmt.Printf("🔧 Task %d finished in %v and sum is %d \n", id, time.Since(start), sum)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}

}
