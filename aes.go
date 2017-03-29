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

func Decrypt(ciphertext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext = PKCS5Unpadding(ciphertext)
	return ciphertext
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
