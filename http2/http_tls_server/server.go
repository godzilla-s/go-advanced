package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//serverCrt = "server.crt"
	//serverKey = "server.key"
	caFile    = "../ca/ca.crt"
	serverCrt = "../ca/server.crt"
	serverKey = "../ca/server.key"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this http server with tls")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server start...")
	server_tls2()
}

func server_tls1() {
	err := http.ListenAndServeTLS(":8081", serverCrt, serverKey, nil)
	if err != nil {
		fmt.Println("listen serve error:", err)
	}
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this is server with tls certificate")
}

func server_tls2() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	pool.AppendCertsFromPEM(caCrt)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert, // 强制校验客户端的证书
			ClientCAs:  pool,
		},
	}
	err = srv.ListenAndServeTLS(serverCrt, serverKey)
	if err != nil {
		log.Fatal("Listen and serve error:", err)
	}
}

// test:
// go run server.go

// curl https://localhost:8081
// 会报错
// curl -k https://localhost:8081
