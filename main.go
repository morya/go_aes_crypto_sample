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

	for _, s := range samples {
		out := Encrypt([]byte(s), key, iv)

		log.Printf("%-30v %X\n", s, out)
	}

}
