package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"go-advanced/crypto/des"
	"math/big"
)

func genKeyPair() (*ecdsa.PrivateKey, ecdsa.PublicKey) {
	prk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	pubkey := prk.PublicKey
	return prk, pubkey
}

// 获取共享密钥
func getdh(prk *ecdsa.PrivateKey, puk ecdsa.PublicKey) *big.Int {
	x, _ := puk.Curve.ScalarMult(puk.X, puk.Y, prk.D.Bytes())
	return x
}

func decEncrypt(prk *ecdsa.PrivateKey, puk ecdsa.PublicKey, plainData []byte) []byte {
	dh := getdh(prk, puk)
	key := dh.Bytes()
	encData, err := des.TriDesEncrypt(plainData, key[:24])
	if err != nil {
		panic(err)
	}
	return encData
}

func decDecrypt(prk *ecdsa.PrivateKey, puk ecdsa.PublicKey, encData []byte) []byte {
	dh := getdh(prk, puk)
	key := dh.Bytes()
	orgData, err := des.TriDesDecrypt(encData, key[:24])
	if err != nil {
		panic(err)
	}
	return orgData
}

func main() {
	prk1, puk1 := genKeyPair()
	prk2, puk2 := genKeyPair()

	//fmt.Println(prk1, puk1)
	//fmt.Println(prk2, puk2)
	//fmt.Println(getdh(prk1, puk2))
	//fmt.Println(getdh(prk2, puk1))
	data := []byte("hello zhangfei")
	encData := decEncrypt(prk1, puk2, data)
	orgData := decDecrypt(prk2, puk1, encData)
	fmt.Println(string(orgData))
}
