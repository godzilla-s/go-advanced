package main

import (
	"flag"
	"fmt"
	"go-advanced/network"
)

var (
	port int
	mode string
)

func main() {
	flag.IntVar(&port, "p", 0, "port")
	flag.StringVar(&mode, "m", "", "server or client")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)
	if mode == "server" {
		network.Server(addr)
	}
	if mode == "client" {
		network.Client(addr)
	}
}
