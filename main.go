// c.go
package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stderr)

	iv := []byte("0123456789ABCDEF")
	key := []byte("dca9b9fd72d0e27b")

	var samples []string = []string{
		"abc",
		"0123456789abcdef",
		"aaabbbeeefffgggaasdfa",
	}

	for _, txt := range samples {
		cipher := Encrypt([]byte(txt), key, iv)
		txt2 := Decrypt(cipher, key, iv)

		log.Printf("txt=%v, txt2=%s, cipher=%X\n", txt, txt2, cipher)
	}

}
