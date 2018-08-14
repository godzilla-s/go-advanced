package main

import (
	"fmt"
	"path/filepath"
	"runtime"
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

func test1() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 有可能发生异常
func select_rand() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

type People2 interface {
	Show()
}

type Student2 struct{}

func (stu *Student2) Show() {}

func live() People2 {
	var stu *Student2
	return stu
}

func fileabs() {
	abspath, err := filepath.Abs("./docker-compose/crypto-config")
	if err != nil {
		fmt.Println("abs path error:", err)
		return
	}
	fmt.Println("abs path:", abspath)
}
func main() {
	//var p People = &Student{}
	//fmt.Println(p.Speak("hello"))

	//def_call()

	// x := new(big.Int)
	// x.SetBytes([]byte("2E87D59EE650BF66AA958228E0F9C9F1C64AE66E"))
	// fmt.Println("%x", x)

	// 或运算的包含使用
	//x := Apple | Peach | Banana
	//fmt.Println(x)
	//fmt.Println(x|Jujube == x, x|Apple == x, x|Peach == x, x|Banana == x)
	//test1()
	// select_rand()

	fileabs()
}

//./peer.sh chaincode instantiate -o orderer.example.com:7050 --tls true --cafile ./tlsca.example.com-cert.pem -C mychannel -n demo -v 0.0.1 -c '{"Args":["init"]}' -P "OR('Org1MSP.member','Org2MSP.member')"
