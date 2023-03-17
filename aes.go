package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

// 加密
func AesEncrypt(encodeStr string, key string, iv string, process func(crypt []byte) []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	// 根据key 生成密文
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypt := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypt, encodeBytes)
	if process != nil {
		crypt = process(crypt)
	}
	return base64.StdEncoding.EncodeToString(crypt), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	// 填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 解密
func AesDecrypt(decodeStr string, key []byte, iv string) ([]byte, error) {
	// 先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	decodeBytes = []byte(strings.Replace(string(decodeBytes), iv, "", 1))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
