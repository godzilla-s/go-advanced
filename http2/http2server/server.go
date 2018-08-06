package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

const (
	_HTTP2URLBase = "https://127.0.0.1:8000"
	_CertFile     = "/Users/zuvakin/workspace/src/go-advanced/http2/pem/cert.pem"
	_KeyFile      = "/Users/zuvakin/workspace/src/go-advanced/http2/pem/key.pem"
)

type handle func(w http.ResponseWriter, r *http.Request)

func main() {
	httpMux, http2Mux := getHttpMux()
	log.Println("server http run ....")
	go httpSrv(httpMux)
	httpsSrv(http2Mux)
}

func getHttpMux() (httpMux, http2Mux *http.ServeMux) {
	httpMux = http.NewServeMux()
	http2Mux = http.NewServeMux()

	x := make(map[string]handle, 0)
	x["/"] = Home
	x["/v1"] = HelloV1

	for k, v := range x {
		redirectURL := http.RedirectHandler(_HTTP2URLBase, 307)
		httpMux.Handle(k, redirectURL)
		http2Mux.HandleFunc(k, v)
	}
	return
}

func httpSrv(mux *http.ServeMux) {
	log.Fatal(http.ListenAndServe(":8001", mux))
}

func httpsSrv(mux *http.ServeMux) {
	srv := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	http2.VerboseLogs = true
	http2.ConfigureServer(srv, &http2.Server{})

	log.Fatal(srv.ListenAndServeTLS(_CertFile, _KeyFile))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Protocal: %s\n", r.Proto)
	fmt.Fprintf(w, "Home")
}

func HelloV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Hello V1")
}
