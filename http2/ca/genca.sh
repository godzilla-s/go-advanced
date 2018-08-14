#!/bin/bash

# ca私钥
openssl genrsa -out ca.key 2048 

sleep 0.5

# ca证书
openssl req -x509 -new -nodes -key ca.key -subj "/CN=brynZhu.com" -days 3650 -out ca.crt

sleep 0.5

# server私钥
openssl genrsa -out server.key 2048 

sleep 0.5 

# server证书请求文件
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr

sleep 0.5

# 用ca私钥签发server的数字证书
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000