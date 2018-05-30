package main

import (
	"fmt"
	"go-advanced/crypto/des"
)

// 长度必须为8
var deskey = []byte("12345678")

// 3des 长度为24
var des3Key = []byte("123456781234567887654321")

func main() {
	data := []byte("hello,this is new message")
	enc, err := des.DesEncrypt(data, deskey)
	if err != nil {
		panic(err)
	}
	fmt.Println(enc)

	dec, err := des.DesDecrypt(enc, deskey)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dec))

	enc, err = des.TriDesEncrypt(data, des3Key)
	if err != nil {
		panic(err)
	}

	fmt.Println(enc)

	dec, err = des.TriDesDecrypt(enc, des3Key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dec))
}
