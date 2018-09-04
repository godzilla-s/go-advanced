package ecc

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestEcdsaGen(t *testing.T) {
	prk, err := GenerateKey()
	if err != nil {
		t.Fatalf("fail to gen:%v", err)
	}

	fmt.Println(prk.Curve)

	fmt.Println("pubkey: ", PubkeyID(prk))
}

func TestPrk2PEM(t *testing.T) {
	prk, err := GenerateKey()
	if err != nil {
		t.Fatal("fail to generate key:", err)
	}

	prkBytes := prk.D.Bytes()
	paddPrk := make([]byte, (prk.Curve.Params().N.BitLen()+7)/8)
	copy(paddPrk[len(paddPrk)-len(prkBytes):], prkBytes)

	//ans1Bytes, err := ans1.Marshal()

	bytes, err := x509.MarshalECPrivateKey(prk)
	if err != nil {
		t.Fatal(err)
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: bytes})
	fmt.Println("pem:", string(pemBytes))
	err = ioutil.WriteFile("./tempPrk.pem", pemBytes, 0700)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPEM2Prk(t *testing.T) {
	bytes, err := ioutil.ReadFile("./tempPub.pem")
	if err != nil {
		t.Fatal(err)
	}

	block, _ := pem.Decode(bytes)

	if x509.IsEncryptedPEMBlock(block) {
		// TODO
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Fatal(err)
	}

	pubKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("not ecdsa public key")
	}
	fmt.Println(pubKey.X, pubKey.Y)
}

func TestPubkey2PEM(t *testing.T) {
	prk, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := x509.MarshalPKIXPublicKey(&prk.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: bytes})
	fmt.Println("pem:", string(pemBytes))
	err = ioutil.WriteFile("./tempPub.pem", pemBytes, 0700)
	if err != nil {
		t.Fatal(err)
	}
}
