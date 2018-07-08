package main

import (
	"flag"
	"fmt"
	"go-advanced/network/chat"
)

var (
	port int
	mode string
	typ  int
	id   string
)

func main() {
	flag.IntVar(&port, "p", 0, "port")
	flag.IntVar(&typ, "t", 0, "type")
	flag.StringVar(&mode, "m", "", "server or client")
	flag.StringVar(&id, "id", "", "id")
	flag.Parse()

	if port == 0 || id == "" /*|| mode == "" || typ == 0 */ {
		panic("invalid argument")
	}

	addr := fmt.Sprintf(":%d", port)
	chatServe(addr, id)
	//network.Run(typ, mode, addr)

}

func chatServe(addr, id string) {
	chat := chat.New(addr, id)
	chat.Start()
}
