package rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func callRPC(rpc *rpc.Client) {
	var desc string
	rpc.Call("Point.Quadrant", Point{2, 5}, &desc)
	fmt.Println("desc: ", desc)

	var length float64
	rpc.Call("Line.Length", Line{A: Point{3, 7}, B: Point{4, 9}}, &length)
	fmt.Println("length: ", length)

	var area float64
	rpc.Call("Rect.Area", Rect{12, 32}, &area)
	fmt.Println("Rect area: ", area)

	rpc.Call("Circle.Area", Circle{R: 12}, &area)
	fmt.Println("circle area: ", area)
}

func TCP_client() {
	rpc, err := rpc.Dial("tcp", ":8010")
	if err != nil {
		log.Fatal("dail err", err)
	}

	callRPC(rpc)
}

func HTTP_client() {
	rpc, err := rpc.DialHTTP("tcp", ":8010")
	if err != nil {
		log.Fatal(err)
	}
	callRPC(rpc)
}

func JSONRpc_client() {
	rpc, err := jsonrpc.Dial("tcp", ":8010")
	if err != nil {
		log.Fatal(err)
	}

	callRPC(rpc)
}
