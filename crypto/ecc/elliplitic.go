package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const pubkeyLen = 64

type PublicID [pubkeyLen]byte

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(curveFunc(), rand.Reader)
}

func curveFunc() elliptic.Curve {
	return elliptic.P256()
}

// 导出一个ID
func PubkeyID(prk *ecdsa.PrivateKey) PublicID {
	puk := prk.PublicKey
	buf := elliptic.Marshal(curveFunc(), puk.X, puk.Y)
	var id PublicID
	copy(id[:], buf[1:])
	return id
}

func (id PublicID) PublicKey() (*ecdsa.PublicKey, error) {
	if len(id) != pubkeyLen {
		return nil, errors.New("invalid publickey")
	}
	half := pubkeyLen / 2
	puk := &ecdsa.PublicKey{Curve: curveFunc(), X: new(big.Int), Y: new(big.Int)}
	puk.X.SetBytes(id[:half])
	puk.Y.SetBytes(id[half:])
	if !puk.Curve.IsOnCurve(puk.X, puk.Y) {
		return nil, errors.New("id is invalid curve point")
	}
	return puk, nil
}

// 测试
func Run() {
	prk, _ := GenerateKey()

	pukId := PubkeyID(prk)

	_, err := pukId.PublicKey()
	if err != nil {
		fmt.Println(err)
	}
}
