package main

import (
	"flag"
	"fmt"
	"go-advanced/network"
)

var (
	port int
	mode string
	typ  int
)

func main() {
	flag.IntVar(&port, "p", 0, "port")
	flag.IntVar(&typ, "t", 0, "type")
	flag.StringVar(&mode, "m", "", "server or client")
	flag.Parse()

	if port == 0 || mode == "" {
		panic("invalid argument")
	}

	addr := fmt.Sprintf(":%d", port)

	network.Run(typ, mode, addr)
}
