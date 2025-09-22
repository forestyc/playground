package main

import (
	"flag"
	"fmt"
	"github.com/forestyc/playground/pkg/encoding/base64"
	"github.com/forestyc/playground/pkg/security/crypto"
)

func main() {
	var text, method, cipher, plain, key, encode string
	var err error

	flag.StringVar(&key, "key", "", "key")
	flag.StringVar(&text, "text", "", "Cipher text/Plain text")
	flag.StringVar(&method, "method", "", "encrypt/decrypt")
	flag.StringVar(&encode, "encoding", "", "base64 encoding, std/url, use url by default")
	flag.Parse()

	if len(text) == 0 || len(method) == 0 {
		flag.Usage()
		return
	}

	if len(key) == 0 {
		key = "f9718298fcae5859"
	}

	bs64 := &base64.Base64{}
	if len(encode) != 0 && encode == "std" {
		bs64 = bs64.WithEncoding(base64.StdEncoding)
	}

	sm4 := crypto.SM4{}
	switch method {
	case "encrypt":
		plain = text
		var cipherByte []byte
		if cipherByte, err = sm4.Encrypt([]byte(plain), []byte(key)); err != nil {
			panic(err)
		}
		cipher = bs64.Encode(cipherByte)
	case "decrypt":
		cipher = text
		var cipherByte, plainByte []byte
		if cipherByte, err = bs64.Decode(cipher); err != nil {
			panic(err)
		}
		if plainByte, err = sm4.Decrypt(cipherByte, []byte(key)); err != nil {
			panic(err)
		}
		plain = string(plainByte)
	default:
		panic("Unkonwn method")
	}
	fmt.Printf("Plain text: [%s]\n", plain)
	fmt.Printf("Cipher text: [%s]\n", cipher)
}
