// aes.go
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

const BLOCK_SIZE = 16

func Encrypt(txt, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	txt = PKCS5Padding(txt)

	block, err = aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, len(txt))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, txt)

	return ciphertext
}

func Decrypt(input, key, iv []byte) []byte {
	return []byte{}
}

func PKCS5Padding(txt []byte) []byte {
	padLen := BLOCK_SIZE - len(txt)%BLOCK_SIZE
	paddingBytes := bytes.Repeat([]byte{byte(padLen)}, padLen)
	txt = append(txt, paddingBytes...)
	return txt
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
