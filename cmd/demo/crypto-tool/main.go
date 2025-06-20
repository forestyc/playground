package main

import (
	"flag"
	"fmt"
	"github.com/forestyc/playground/pkg/security/crypto"
)

func main() {
	var text, method, cipher, plain, key string
	var err error

	flag.StringVar(&key, "key", "", "key")
	flag.StringVar(&text, "text", "", "Cipher text/Plain text")
	flag.StringVar(&method, "method", "", "encrypt/decrypt")
	flag.Parse()

	if len(text) == 0 || len(method) == 0 {
		flag.Usage()
		return
	}

	if len(key) == 0 {
		key = "f9718298fcae5859"
	}

	sm4 := crypto.SM4{}
	switch method {
	case "encrypt":
		plain = text
		if cipher, err = sm4.EncryptWithBase64([]byte(plain), []byte(key)); err != nil {
			panic(err)
		}
	case "decrypt":
		cipher = text
		var plainByte []byte
		if plainByte, err = sm4.DecryptWithBase64(cipher, []byte(key)); err != nil {
			panic(err)
		}
		plain = string(plainByte)
	default:
		panic("Unkonwn method")
	}
	fmt.Printf("Plain text: [%s]\n", plain)
	fmt.Printf("Cipher text: [%s]\n", cipher)
}
