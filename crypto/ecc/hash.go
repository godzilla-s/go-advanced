package crypto

import "crypto/sha256"

type Hash [32]byte

func toHash(data []byte) Hash {
	d := sha256.New()
	d.Write(data)

	h := d.Sum(nil)
	var hash Hash
	copy(hash[:], h[:])
	return hash
}
