package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
)

type Scanner struct {
	msg  chan string
	scan *bufio.Scanner
}

func (s *Scanner) Scan() {
	for s.scan.Scan() {
		s.msg <- s.scan.Text()
	}
}

func (s *Scanner) Read() <-chan string {
	return s.msg
}

func main() {
	scanner := &Scanner{
		msg:  make(chan string, 5),
		scan: bufio.NewScanner(os.Stdin),
	}
	go scanner.Scan()
	for v := range scanner.Read() {
		fmt.Println("read:", v)
	}
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func wg() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
