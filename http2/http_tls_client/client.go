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
	req_url    = "https://localhost:8081"
	caCertPath = "../ca/ca.crt"
)

func main() {
	client_request3()
}

// 不校验证书
func client_request1() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // client将不会对服务端证书进行验证
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Get(req_url)
	if err != nil {
		fmt.Println("get err:", err)
		return
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// 确保 server服务启动
// go run client.go

// 校验证书
func client_request2() {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("readfile err:", err)
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}

	client := &http.Client{Transport: transport}

	resp, err := client.Get(req_url)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 双向校验证书:
// 		服务端对客户端进行校验
// 		客户端对服务端也进行校验
// ca需要给客户端签发证书

const (
	Client_KEY = "client.key"
	Client_CRT = "client.crt"
)

func client_request3() {
	fmt.Println("双向证书验证")
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatal("read file error:", err)
		return
	}
	// 加入CA证书
	pool.AppendCertsFromPEM(caCrt)

	// 加载客户端证书
	cliCrt, err := tls.LoadX509KeyPair(Client_CRT, Client_KEY)
	if err != nil {
		log.Fatal("local ca file:", err)
		return
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt}, // 客户端证书
		},
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Get(req_url)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
