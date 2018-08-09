package main

import (
	"encoding/base64"
	"fmt"
	"go-advanced/crypto/des"
)

// 长度必须为8
var deskey = []byte("12345678")

// 3des 长度为24
var des3Key = []byte("123456781234567887654321")

func simple() {
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

func main() {
	decrypto_01(encrypto_01("smile2018", ""))
	decrypto_01(encrypto_01("abcd001", ""))
	decrypto_01(encrypto_01("qrtyoerwuy643564356", ""))
	decrypto_01(encrypto_01("bgg3452345345234^&$&%$", ""))
}

// des 对称密钥 k1
// a: 密码明文 -> des加密 -> base64转化 -> 密文密码
func encrypto_01(privKey, data string) string {
	// 先用des加密密钥
	enc, err := des.DesEncrypt([]byte(privKey), deskey)
	if err != nil {
		panic(err)
	}

	// 再用base64转化
	baseStr := base64.StdEncoding.EncodeToString(enc)
	fmt.Println("encode password:", baseStr)
	return baseStr
}

func decrypto_01(encrypt string) {
	enc, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println("deode string error:", err)
		return
	}

	dec, err := des.DesDecrypt(enc, deskey)
	if err != nil {
		fmt.Println("des decrypto error:", err)
		return
	}

	fmt.Println("org:", string(dec))
}
