package main

import (
	"flag"
	"fmt"
	_ "go-advanced/reflect"
	"go-advanced/run"
)

var fname string

func main() {
	flag.StringVar(&fname, "f", "", "function name")
	flag.Parse()

	functions := run.GetFunctions()

	if f, ok := functions[fname]; ok {
		f()
	} else {
		fmt.Printf("function %s not register\n", fname)
	}
}
