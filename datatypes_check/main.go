package main

import (
	"datatype_skill/chanpractice"
	stringManipulation "datatype_skill/string"
	"sync"

	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
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
}
