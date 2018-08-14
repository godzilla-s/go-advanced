#!/bin/bash

# 签发客户端证书

CLI_KEY="client.key"
CLI_CRT="client.crt"
CA_HOME=`cd ../ca;pwd`
CA_CRT="$CA_HOME/ca.crt"
CA_KEY="$CA_HOME/ca.key"

function create()
{
    # 客户端私钥
    openssl genrsa -out $CLI_KEY 2048

    sleep 0.5

    openssl req -new -key $CLI_KEY -subj "/CN=brynZhu_cn" -out client.csr 

    sleep 0.5 

    # ca给client签发证书
    openssl x509 -req -in client.csr -CA $CA_CRT -CAkey $CA_KEY -CAcreateserial -out $CLI_CRT -days 5000
}

if [ "$1" = "clean" ]; then 
    rm -f $CLI_CRT $CLI_KEY client.csr
else
    echo "$CA_CRT"
    create
fi
