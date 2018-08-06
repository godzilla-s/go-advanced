package main

import (
	"fmt"
)

type People interface {
	Speak(s string) string
}

type Student struct{}

func (stu *Student) Speak(s string) string {
	if s == "hello" {
		return "You are a new one"
	} else {
		return "reponse back:" + s
	}
}

func def_call() {
	defer func() { fmt.Println("defer call 01") }()
	defer func() { fmt.Println("defer call 02") }()
	defer func() { fmt.Println("defer call 03") }()

	panic("def call")
}

const (
	Apple = 1 << iota
	Banana
	Peach
	Jujube
)

func main() {
	//var p People = &Student{}
	//fmt.Println(p.Speak("hello"))

	//def_call()

	// x := new(big.Int)
	// x.SetBytes([]byte("2E87D59EE650BF66AA958228E0F9C9F1C64AE66E"))
	// fmt.Println("%x", x)

	// 或运算的包含使用
	x := Apple | Peach | Banana
	fmt.Println(x)
	fmt.Println(x|Jujube == x, x|Apple == x, x|Peach == x, x|Banana == x)
}
