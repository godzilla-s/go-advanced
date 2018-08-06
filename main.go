package main

import (
	"fmt"
	"net"
	"net/http"
)

var s = `curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://7a038b44.m.daocloud.io`

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, s)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8001", nil)
	//go ListenUDP("192.168.1.100:6001")
	//ListenTCP("192.168.1.100:6001")
}

func ListenUDP(addr string) {
	laddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	lsn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("listen udp ...")
	buf := make([]byte, 512)
	for {
		n, _, err := lsn.ReadFromUDP(buf)
		if err != nil {
			return
		}
		fmt.Println("udp read:", buf[:n])
	}
}

func ListenTCP(addr string) {
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}

	lsn, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("listen tcp ...")
	for {
		_, err := lsn.AcceptTCP()
		if err != nil {
			return
		}

	}
}
