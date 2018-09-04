package aes

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	orgData := []byte("abcdejashfrvlbhabclaib45adf33")
	key := []byte("woshikey20181234") // AES key要求16(AES-128)，24(AES-196)，32(AES-256)长度

	enc, err := AesEncrypt(orgData, key)
	if err != nil {
		t.Fatal("encrypt: ", err)
	}
	fmt.Println("enc:", enc, "org len:", len(orgData), "enc len:", len(enc))

	dec, err := AesDecrypt(enc, key)
	if err != nil {
		t.Fatal("dec:", err)
	}
	fmt.Println("dec:", string(dec))

}
