package main

import (
	"flag"
	"go-advanced/grpc/test"
)

var run string

func main() {
	flag.StringVar(&run, "r", "", "run mode")
	flag.Parse()

	if run == "server" {
		test.Server()
	} else if run == "client" {
		test.Client()
	} else {
		panic("invalid args")
	}

}
