package main

import (
	"flag"
	"go-advanced/rpc"
)

var startup string
var mode string

func main() {
	flag.StringVar(&startup, "s", "", "start server or client")
	flag.StringVar(&mode, "t", "", "tcp, http, jsonrpc")
	flag.Parse()

	if startup == "server" {
		switch mode {
		case "tcp":
			rpc.TCP_server()
		case "http":
			rpc.HTTP_server()
		case "jsonrpc":
			rpc.JSONRpc_server()
		}
	}

	if startup == "client" {
		switch mode {
		case "tcp":
			rpc.TCP_client()
		case "http":
			rpc.HTTP_client()
		case "jsonrpc":
			rpc.JSONRpc_client()
		}
	}
}
