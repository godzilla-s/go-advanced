package rpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func register_service() {
	point := new(Point)
	line := new(Line)
	rect := new(Rect)
	circle := new(Circle)

	rpc.Register(point)
	rpc.Register(line)
	rpc.Register(rect)
	rpc.Register(circle)
}

func TCP_server() {
	register_service()

	laddr, err := net.ResolveTCPAddr("tcp4", ":8010")
	if err != nil {
		log.Println("resolve tcp error:", err)
		return
	}

	lsn, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		conn, err := lsn.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(conn)
	}
}

func HTTP_server() {
	register_service()

	rpc.HandleHTTP()

	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatalf("listen error:%v", err)
	}
}

func JSONRpc_server() {
	register_service()

	laddr, err := net.ResolveTCPAddr("tcp4", ":8010")
	if err != nil {
		log.Println("resolve tcp error:", err)
		return
	}

	lsn, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		conn, err := lsn.Accept()
		if err != nil {
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
