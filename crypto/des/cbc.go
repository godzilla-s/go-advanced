// CBC 加密模式
package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// PKCS5Padding 填充数据
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS5UnPadding 还原数据
func PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unpaddind := int(data[length-1])
	return data[:(length - unpaddind)]
}

// DesEncript DES加密
// origData 加密数据， key 秘钥
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	origData = PKCS5Padding(origData, block.BlockSize())
	// 调用加密函数
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 调用解密函数
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// TriDesEncrypt 3DES加密
func TriDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// TriDesDecrypt 3DES解密
func TriDesDecrypt(encData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(encData))
	blockMode.CryptBlocks(origData, encData)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
