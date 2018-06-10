package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func GenerateRSAKey(bits int) {
	prk, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}

	decData := x509.MarshalPKCS1PrivateKey(prk)
	block := &pem.Block{
		Type:  "Private",
		Bytes: decData,
	}

	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}

	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
	}
	return
}

// Encrypt 加密
func Encrypt(prk []byte, orig []byte) ([]byte, error) {
	block, _ := pem.Decode(prk)
	if block == nil {
		return nil, fmt.Errorf("decode error")
	}

	pubInf, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInf.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, orig)
}

func Decrypt(prk, cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(prk)
	if block == nil {
		return nil, errors.New("get private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
